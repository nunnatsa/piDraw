package hat

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
	"github.com/nathany/bobblehat/sense/stick"
	"github.com/nunnatsa/piDraw/datatype"
)

const (
	rmask color.Color = 0b1111100000000000
	gmask color.Color = 0b0000011111100000
	bmask color.Color = 0b0000000000011111
)

type Hat struct {
	events chan<- datatype.HatEvent
	screen <-chan *datatype.DisplayMessage
	input  *stick.Device
}

func NewHat(e chan<- datatype.HatEvent, s <-chan *datatype.DisplayMessage) *Hat {
	h := &Hat{
		events: e,
		screen: s,
	}

	go h.do()
	return h
}

func (h *Hat) init() {
	var err error
	h.input, err = stick.Open("/dev/input/event0")
	if err != nil {
		log.Panic(err)
	}
	screen.Clear()
}

func (h *Hat) do() {
	h.init()
	// Set up a signals channel (stop the loop using Ctrl-C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	defer log.Printf("HAT even handling crached")
	for {
		select {
		case <-signals:
			screen.Clear()
			fmt.Println("")
			os.Exit(0)
		case event := <-h.input.Events:
			switch event.Code {
			case stick.Enter:
				h.events <- datatype.Pressed
				log.Println("Joystick Event: Pressed")
			case stick.Up:
				h.events <- datatype.MoveUp
				log.Println("Joystick Event: MoveUp")
			case stick.Down:
				h.events <- datatype.MoveDown
				log.Println("Joystick Event: MoveDown")
			case stick.Left:
				h.events <- datatype.MoveLeft
				log.Println("Joystick Event: MoveLeft")
			case stick.Right:
				h.events <- datatype.MoveRight
				log.Println("Joystick Event: MoveRight")
			}
		case screenChange := <-h.screen:
			//log.Println("HAT: got screen update:", screenChange)
			
			h.drawScreen(screenChange)
		}
	}
}

func (h *Hat) drawScreen(screenChange *datatype.DisplayMessage) {
	fb := screen.NewFrameBuffer()
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			screenPixel := screenChange.Screen[y][x]
			hatPixel := toHatColor(screenPixel)
			fb.SetPixel(x, y, hatPixel)
		}
	}
	fb.SetPixel(int(screenChange.CursorX), int(screenChange.CursorY), rmask)
	err := screen.Draw(fb)
	if err != nil {
		log.Println("error while printing to HAT display:", err)
	}
}

func toHatColor(c datatype.Color) color.Color {
	r := color.Color(c>>8) & rmask
	g := color.Color(c>>3) & gmask
	b := color.Color(c) & bmask

	return r | g | b
}
