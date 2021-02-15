package webapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nunnatsa/piDraw/datatype"
	"github.com/nunnatsa/piDraw/notifier"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultPixelSize = 5
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1500,
		WriteBufferSize: 1500,
	}
)

type colorMessage struct {
	Color datatype.Color `json:"color"`
}

type ClientActions struct {
	mux     *http.ServeMux
	mailbox *notifier.Notifier
	ch      chan<- datatype.ClientEvent
}

func (ca ClientActions) GetMux() *http.ServeMux {
	return ca.mux
}

func NewClientAction(mailbox *notifier.Notifier, port uint16, ch chan<- datatype.ClientEvent) *ClientActions {
	mux := http.NewServeMux()
	ca := &ClientActions{mux: mux, mailbox: mailbox, ch: ch}
	mux.Handle("/", newIndexPage(port))
	mux.HandleFunc("/api/canvas/color", ca.setColor)
	mux.HandleFunc("/api/canvas/reset", ca.reset)
	mux.HandleFunc("/api/canvas/register", ca.register)
	mux.HandleFunc("/api/canvas/download", ca.downloadImage)

	return ca
}

func (ca ClientActions) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
		}

		defer conn.Close()

		subscription := make(chan []byte)

		id := ca.mailbox.Subscribe(subscription)
		defer ca.mailbox.Unsubscribe(id)
		ca.ch <- datatype.ClientEvent{EventType: datatype.EventClientRegistered, Data: id}

		for js := range subscription {
			log.Printf("got event; updating the %d\n", id)
			if err := conn.WriteMessage(websocket.TextMessage, js); err != nil {
				log.Printf("failed to send message to the client %d: %v\n", id, err)
				return
			}
		}
		log.Println("Connection is closed")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ca ClientActions) reset(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ca.ch <- datatype.ClientEvent{EventType: datatype.EventTypeReset}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ca ClientActions) setColor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		enc := json.NewDecoder(r.Body)
		color := &colorMessage{}
		if err := enc.Decode(color); err != nil {
			log.Println("error while handling POST /api/Canvas/color", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "can't process body"}`)
			return
		}
		ca.ch <- datatype.ClientEvent{
			EventType: datatype.EventTypeColorChange,
			Data:      color.Color,
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ca ClientActions) downloadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		imageChannel := make(chan [][]datatype.Color)
		_ = r.ParseForm()

		pixelSizeStr := r.Form.Get("pixelSize")
		pixelSize, err := strconv.Atoi(pixelSizeStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "pixelSize must be a number")
			return
		}

		if pixelSize < 1 || pixelSize > 20 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "pixelSize must between 1 to 10")
			return
		}

		ca.ch <- datatype.ClientEvent{
			EventType: datatype.EventTypeDownloadRequest,
			Data:      imageChannel,
		}

		done := make(chan bool)
		go func(ch chan<- [][]datatype.Color) {
			select {
			case imageBytes := <-imageChannel:
				createPng(w, imageBytes, pixelSize)
			case <-time.After(time.Second * 3):
				w.WriteHeader(http.StatusServiceUnavailable)
			}
			close(done)
		}(imageChannel)

		<-done
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func createPng(w http.ResponseWriter, pixels [][]datatype.Color, pixelSize int) {
	if len(pixels) == 0 {
		return
	}

	height := len(pixels)
	width := len(pixels[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width * pixelSize, height * pixelSize}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			setPixel(img, x, y, pixels[y][x], pixelSize)
		}
	}

	w.Header().Add("Content-Disposition", `attachment; filename="untitled.jpg"`)
	w.Header().Add("Context-Type", "image/png")
	// Encode as PNG.
	png.Encode(w, img)
}

func setPixel(img *image.RGBA, x int, y int, pixel datatype.Color, pixelSize int) {
	x = x*pixelSize
	y = y*pixelSize
	for x1 := x; x1 < x + pixelSize; x1++ {
		for y1 := y; y1 < y+pixelSize; y1++ {
			img.Set(x1, y1, colorToImageColor(pixel))
		}
	}
}

func colorToImageColor(c datatype.Color) color.RGBA {
	R := uint8((c >> 16) & 0xFF)
	G := uint8((c >> 8) & 0xFF)
	B := uint8(c & 0xFF)
	return color.RGBA{A: 0xFF, R: R, G: G, B: B}
}
