package canvas

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/nunnatsa/piDraw/datatype"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1500,
		WriteBufferSize: 1500,
	}
)

// Board is the draing board.
type Board struct {
	Canvas *Canvas `json:"canvas,omitempty"`
	Cursor *Cursor `json:"cursor,omitempty"`
	Window *Window `json:"window,omitempty"`
	reg    *Notifier
	events <-chan datatype.HatEvent
	screen chan<- *datatype.DisplayMessage
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
		Window: c.prepareWindow(windowSize, windowSize),
		reg: &Notifier{
			clientMap: make(map[int64]chan []byte),
			idp:       &idProvider{lock: &sync.Mutex{}},
		},
		events: events,
		screen: screen,
	}

	go b.do()

	return b
}

func (b *Board) do () {
	for event := range b.events {
		switch event {
		case datatype.Pressed:
			b.DrawPixel()
		case datatype.MoveUp:
			b.MoveUp()
		case datatype.MoveDown:
			b.MoveDown()
		case datatype.MoveLeft:
			b.MoveLeft()
		case datatype.MoveRight:
			b.MoveRight()
		}

		b.Update()
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
	b.Canvas = newCanvas()
	b.Cursor = &Cursor{
		X:     centerX,
		Y:     centerY,
		Color: 0xFFFFFF,
	}
	b.Window = b.Canvas.prepareWindow(windowSize, windowSize)
}

func (b *Board) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			log.Println("Error:", err)
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
		}

		ch := make(chan []byte)

		id := b.reg.Subscribe(ch)
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

		for range ch {
			if err := conn.WriteMessage(websocket.TextMessage, js); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (b *Board) Update() {
	msg := datatype.NewDisplayMessage(b.Window.matrix, b.Cursor.X, b.Cursor.Y)
	b.screen <- msg

	js, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	} else {
		b.reg.Notify(js)
	}
}
