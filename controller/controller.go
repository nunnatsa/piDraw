package controller

import (
	"encoding/json"
	"github.com/nunnatsa/piDraw/canvas"
	"github.com/nunnatsa/piDraw/datatype"
	"github.com/nunnatsa/piDraw/hat"
	"github.com/nunnatsa/piDraw/notifier"
	"log"
	"runtime"
)

type Controller struct {
	board        *canvas.Board
	theHat       *hat.Hat
	hatEvents    chan datatype.HatEvent
	screenEvents chan *datatype.DisplayMessage
	clientEvents chan datatype.ClientEvent
	mailbox      *notifier.Notifier
}

func NewController(width, height uint8, mailbox *notifier.Notifier) *Controller {
	hatEvents := make(chan datatype.HatEvent)
	screenEvents := make(chan *datatype.DisplayMessage)
	clientEvents := make(chan datatype.ClientEvent)
	board := canvas.NewBoard(width, height)

	c := &Controller{
		board:        board,
		hatEvents:    hatEvents,
		screenEvents: screenEvents,
		clientEvents: clientEvents,
		mailbox:      mailbox,
	}

	if "arm" == runtime.GOARCH {
		c.theHat = hat.NewHat(hatEvents, screenEvents)
	} else {
		hat.NewHatMock(hatEvents, screenEvents)
	}

	return c
}

func (c Controller) Start() {
	go c.do()
}

func (c Controller) GetClientEvents() chan<- datatype.ClientEvent {
	return c.clientEvents
}

func (c *Controller) do() {
	c.updateScreen()
	for {
		changed := false
		select {
		case event := <-c.hatEvents:
			switch event {
			case datatype.Pressed:
				log.Println("HAT Event: Pressed")
				c.board.DrawPixel()
				changed = true
			case datatype.MoveUp:
				log.Println("HAT Event: MoveUp")
				c.board.MoveUp()
				changed = true
			case datatype.MoveDown:
				log.Println("HAT Event: MoveDown")
				c.board.MoveDown()
				changed = true
			case datatype.MoveLeft:
				log.Println("HAT Event: MoveLeft")
				c.board.MoveLeft()
				changed = true
			case datatype.MoveRight:
				log.Println("HAT Event: MoveRight")
				c.board.MoveRight()
				changed = true
			}
		case event := <-c.clientEvents:
			switch event.EventType {
			case datatype.EventTypeReset:
				c.board.Reset()
				changed = true
			case datatype.EventTypeColorChange:
				c.board.Cursor.SetColor(event.Data.(datatype.Color))
				changed = true
			case datatype.EventClientRegistered:
				c.registered(event.Data.(int64))
			case datatype.EventTypeDownloadRequest:
				c.getImagePixels(event.Data.(chan [][]datatype.Color))
			}
		}
		if changed {
			c.Update()
		}
	}

}

func (c *Controller) Update() {
	c.updateScreen()

	js, err := json.Marshal(c.board)
	if err != nil {
		log.Println(err)
	} else {
		c.mailbox.Notify(js)
	}
}

func (c *Controller) updateScreen() {
	msg := datatype.NewDisplayMessage(c.board.Window.Matrix, c.board.Cursor.X-c.board.Window.X, c.board.Cursor.Y-c.board.Window.Y)
	c.screenEvents <- msg
}

func (c *Controller) registered(id int64) {
	js, err := json.Marshal(c.board)
	if err != nil {
		log.Println(err)
	} else {
		c.mailbox.NotifyOne(id, js)
	}
}

func (c *Controller) getImagePixels(ch chan<- [][]datatype.Color) {
	pixels := make([][]datatype.Color, len(c.board.Canvas))
	for i, line := range c.board.Canvas {
		pixels[i] = make([]datatype.Color, len(line))
		copy(pixels[i], line)
	}

	ch <- pixels
	close(ch)
}
