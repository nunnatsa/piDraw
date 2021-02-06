package frontend

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var Mux *http.ServeMux = http.NewServeMux()

func init() {
	t, err := template.ParseFiles("static/index.gohtml")
	if err != nil {
		log.Panic(err)
	}
	f, err := os.OpenFile("static/index.html", os.O_CREATE | os.O_RDWR | os.O_TRUNC, 0664)
	if err != nil {
		log.Panic(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Panic(err)
	}
	err = t.Execute(f, hostname)
	if err != nil {
		log.Panic(err)
	}

	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
}
