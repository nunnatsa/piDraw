package canvas

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nunnatsa/piDraw/datatype"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1500,
		WriteBufferSize: 1500,
	}
)

// Board is the draing board.
type Board struct {
	Canvas       *Canvas `json:"canvas,omitempty"`
	Cursor       *Cursor `json:"cursor,omitempty"`
	Window       *Window `json:"window,omitempty"`
	reg          *Notifier
	hatEvents    <-chan datatype.HatEvent
	clientEvents chan clientEvent
	screen       chan<- *datatype.DisplayMessage
}

type colorMessage struct {
	Color datatype.Color `json:"color"`
}

type clientEventType int32

const (
	eventTypeReset clientEventType = iota
	eventTypeColorChange
)

type clientEvent struct {
	eventType clientEventType
	data      interface{}
}

// NewBoard initiate a new Board
func NewBoard(events <-chan datatype.HatEvent, screen chan<- *datatype.DisplayMessage) *Board {
	c := newCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{
			X:     centerX,
			Y:     centerY,
			Color: 0xFFFFFF,
		},
		Window:       c.prepareWindow(windowSize, windowSize),
		reg:          newNotifier(),
		hatEvents:    events,
		clientEvents: make(chan clientEvent),
		screen:       screen,
	}

	go b.do()

	return b
}

func (b *Board) do() {
	for {
		changed := false
		select {
		case event := <-b.hatEvents:
			switch event {
			case datatype.Pressed:
				log.Println("HAT Event: Pressed")
				b.DrawPixel()
				changed = true
			case datatype.MoveUp:
				log.Println("HAT Event: MoveUp")
				b.MoveUp()
				changed = true
			case datatype.MoveDown:
				log.Println("HAT Event: MoveDown")
				b.MoveDown()
				changed = true
			case datatype.MoveLeft:
				log.Println("HAT Event: MoveLeft")
				b.MoveLeft()
				changed = true
			case datatype.MoveRight:
				log.Println("HAT Event: MoveRight")
				b.MoveRight()
				changed = true
			}
		case event := <-b.clientEvents:
			switch event.eventType {
			case eventTypeReset:
				b.Reset()
				changed = true
			case eventTypeColorChange:
				b.Cursor.SetColor(event.data.(datatype.Color))
				changed = true
			}
		}
		if changed {
			b.Update()
		}
	}
}

// MoveUp moves the cursor one pixel up, if not already on top
func (b *Board) MoveUp() {
	b.Cursor.MoveUp()
	if b.Cursor.Y < b.Window.Y {
		b.Window.Y = b.Cursor.Y
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// MoveDown moves the cursor one pixel down, if not already on bottom
func (b *Board) MoveDown() {
	b.Cursor.MoveDown()
	if b.Cursor.Y > b.Window.Y+windowSize-1 {
		b.Window.Y++
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// MoveLeft moves the cursor one pixel down, if not already on bottom
func (b *Board) MoveLeft() {
	b.Cursor.MoveLeft()
	if b.Cursor.X < b.Window.X {
		b.Window.X = b.Cursor.X
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// MoveRight moves the cursor one pixel down, if not already on bottom
func (b *Board) MoveRight() {
	b.Cursor.MoveRight()
	if b.Cursor.X > b.Window.X+windowSize-1 {
		b.Window.X++
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// DrawPixel writes the pixel the the cursor is currently pointing at, to the Color of the cursor
func (b *Board) DrawPixel() {
	b.Canvas.Set(b.Cursor)
}

// DeletePixel writes the pixel the the cursor is currently pointing at, to the Color of the cursor
func (b *Board) DeletePixel() {
	b.Canvas.Delete(b.Cursor)
}

// Reset return the board to the initiate state
func (b *Board) Reset() {
	log.Println("Reseting the canvas")
	b.Canvas = newCanvas()
	b.Cursor = &Cursor{
		X:     centerX,
		Y:     centerY,
		Color: 0xFFFFFF,
	}
	b.Window = b.Canvas.prepareWindow(windowSize, windowSize)
	log.Println("canvas is clean now")
}

func (b *Board) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
		}

		defer conn.Close()

		subscription := make(chan []byte)

		id := b.reg.Subscribe(subscription)
		defer b.reg.Unsubscribe(id)

		js, err := json.Marshal(b)
		if err != nil {
			log.Println(err)
		} else {
			if err := conn.WriteMessage(websocket.TextMessage, js); err != nil {
				log.Println(err)
				return
			}
		}

		for js := range subscription {
			log.Printf("got event; updating the %d\n", id)
			if err := conn.WriteMessage(websocket.TextMessage, js); err != nil {
				log.Printf("failed to send message to the client %d: %v\n", id, err)
				return
			}
		}
		log.Println("Connection is closed")
	} else if r.URL.Path == "/reset" && r.Method == http.MethodPost {
		b.clientEvents <- clientEvent{eventType: eventTypeReset}
	} else if r.URL.Path == "/color" && r.Method == http.MethodPost {
		enc := json.NewDecoder(r.Body)
		color := &colorMessage{}
		if err := enc.Decode(color); err != nil {
			log.Println("error while handling POST /api/canvas/color", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "can't process body"}`)
			return
		}
		b.clientEvents <- clientEvent{eventType: eventTypeColorChange, data: color.Color}
	}
}

func (b *Board) Update() {
	msg := datatype.NewDisplayMessage(b.Window.matrix, b.Cursor.X-b.Window.X, b.Cursor.Y-b.Window.Y)
	b.screen <- msg

	js, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	} else {
		b.reg.Notify(js)
	}
}
