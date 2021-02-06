package main

import (
	"github.com/nunnatsa/piDraw/canvas"
	"github.com/nunnatsa/piDraw/datatype"
	"log"
	"net/http"

	"github.com/nunnatsa/piDraw/frontend"
)

func main() {

	events := make(chan datatype.HatEvent)
	screen := make(chan *datatype.DisplayMessage)
	c := canvas.NewBoard(events, screen)

	frontend.Mux.Handle("/api/canvas", c)
	log.Panic(http.ListenAndServe(":8080", frontend.Mux))
}
