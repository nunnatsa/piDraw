package webapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nunnatsa/piDraw/datatype"
	"github.com/nunnatsa/piDraw/notifier"
	"log"
	"net/http"
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
