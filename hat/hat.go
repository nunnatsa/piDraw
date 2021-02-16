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

// The format of the HAT color is 16-bit: 5 MS bits are the red color, the middle 6 bits are
// green and the 5 LB bits are blue
// rrrrrggggggbbbbb
const (
	redColor color.Color    = 0b1111100000000000
	rmask    datatype.Color = 0b111110000000000000000000
	gmask    datatype.Color = 0b000000001111110000000000
	bmask    datatype.Color = 0b000000000000000011111000
)

// to convert 24-bit color to 16-bit color, we are taking only the 5 (for red and
// blue) or 6 (for green) MS bits
func toHatColor(c datatype.Color) color.Color {
	r := color.Color((c & rmask) >> 8)
	g := color.Color((c & gmask) >> 5)
	b := color.Color((c & bmask) >> 3)

	return r | g | b
}

type HATInterface interface {
	SetChannels(e chan<- datatype.HatEvent, s <-chan *datatype.DisplayMessage)
	Start()
}

type Hat struct {
	events chan<- datatype.HatEvent
	screen <-chan *datatype.DisplayMessage
	input  *stick.Device
}

func (h *Hat) SetChannels(e chan<- datatype.HatEvent, s <-chan *datatype.DisplayMessage) {
	h.events = e
	h.screen = s
}

func (h *Hat) Start() {
	go h.do()
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
	fb.SetPixel(int(screenChange.CursorX), int(screenChange.CursorY), redColor)
	err := screen.Draw(fb)
	if err != nil {
		log.Println("error while printing to HAT display:", err)
	}
}
