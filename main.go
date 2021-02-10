package main

import (
	"log"
	"net/http"

	"github.com/nunnatsa/piDraw/hat"

	"github.com/nunnatsa/piDraw/canvas"
	"github.com/nunnatsa/piDraw/datatype"
	"github.com/nunnatsa/piDraw/frontend"
)

func main() {

	events := make(chan datatype.HatEvent)
	screen := make(chan *datatype.DisplayMessage, 100)
	c := canvas.NewBoard(events, screen)
	_ = hat.NewHat(events, screen)

	frontend.Mux.Handle("/api/canvas/", http.StripPrefix("/api/canvas", c))
	log.Panic(http.ListenAndServe(":8080", frontend.Mux))
}
