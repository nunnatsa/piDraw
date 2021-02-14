package main

import (
	"flag"
	"fmt"
	"github.com/nunnatsa/piDraw/hat"
	"github.com/nunnatsa/piDraw/webapp"
	"log"
	"net/http"

	"github.com/nunnatsa/piDraw/canvas"
	"github.com/nunnatsa/piDraw/datatype"
)

var (
	port   uint
	height uint
	width  uint
)

func init() {
	flag.UintVar(&port, "port", 8080, "webapp port")
	flag.UintVar(&height, "height", 3, "canvas height in multiplies of 8 ( e.g. for height of 24, set to 3")
	flag.UintVar(&width, "width", 3, "canvas width in multiplies of 8 ( e.g. for width of 24, set to 3")

	flag.Parse()

	if port > 0xFFFF {
		panic(fmt.Sprintf("port number can't be more than %d", 0xFFFF))
	}
	if height == 0 || height > 5 {
		panic(fmt.Sprintln("height must be from 1 to 5"))
	}
	if width == 0 || width > 5 {
		panic(fmt.Sprintln("width must be from 1 to 5"))
	}
}

func main() {
	events := make(chan datatype.HatEvent)
	screen := make(chan *datatype.DisplayMessage, 100)

	c := canvas.NewBoard(events, screen, uint8(width), uint8(height))
	_ = hat.NewHat(events, screen)
	mux := webapp.GetMux(uint16(port))
	mux.Handle("/api/canvas/", http.StripPrefix("/api/canvas", c))
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
