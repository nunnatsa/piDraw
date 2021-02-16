package hat

import (
	term "github.com/nsf/termbox-go"
	"github.com/nunnatsa/piDraw/datatype"
	"log"
	"sync"
)

type HatMock struct {
	e           chan<- datatype.HatEvent
	s           <-chan *datatype.DisplayMessage
	lastMessage *datatype.DisplayMessage
	ready       bool
	lock        *sync.Mutex
}

func NewMock() *HatMock {
	return &HatMock{
		lock: &sync.Mutex{},
	}
}

func (hm *HatMock) SetChannels(e chan<- datatype.HatEvent, s <-chan *datatype.DisplayMessage) {
	hm.e = e
	hm.s = s
}

func (hm *HatMock) GetLastMsg() *datatype.DisplayMessage {
	for {
		hm.lock.Lock()
		var msg *datatype.DisplayMessage
		if hm.ready {
			msg = hm.lastMessage
			hm.ready = false
		}
		hm.lock.Unlock()
		if msg != nil {
			return msg
		}
	}
}

func (hm *HatMock) ConsumeDisplayMessages() {
	for lastMessage := range hm.s {
		hm.lock.Lock()
		hm.ready = true
		hm.lastMessage = lastMessage
		hm.lock.Unlock()
		log.Println("[HAT Mock] Got display event")
	}
}

func (hm *HatMock) Close() {
	term.Flush()
	term.Close()
}

func (hm *HatMock) GoUp() {
	hm.e <- datatype.MoveUp
}
func (hm *HatMock) GoDown() {
	hm.e <- datatype.MoveDown
}
func (hm *HatMock) GoLeft() {
	hm.e <- datatype.MoveLeft
}
func (hm *HatMock) GoRight() {
	hm.e <- datatype.MoveRight
}
func (hm *HatMock) Press() {
	hm.e <- datatype.Pressed
}

func reset() {
	term.Sync() // cosmestic purpose
}

func (hm *HatMock) Start() {
	go hm.ConsumeDisplayMessages()

	go func() {
		defer term.Close()

		if err := term.Init(); err != nil {
			log.Println("Can't get user inputs")
			return
		}

		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyArrowUp:
					hm.e <- datatype.MoveUp
					reset()
				case term.KeyArrowLeft:
					hm.e <- datatype.MoveLeft
					reset()
				case term.KeyArrowDown:
					hm.e <- datatype.MoveDown
					reset()
				case term.KeyArrowRight:
					hm.e <- datatype.MoveRight
					reset()
				case term.KeyEnter:
					hm.e <- datatype.Pressed
					reset()
				case term.KeyCtrlC:
					reset()
					term.Close()
					return
				default:
					reset()
					log.Printf("key pressed: %v (%v)", ev.Key, ev.Ch)
				}
			case term.EventError:
				log.Println("Can't get user inputs")
				return
			}
		}
	}()
}
