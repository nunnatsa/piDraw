package webapp

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
)

type index struct {
	indexPageBytes []byte
}

type indexPageParams struct {
	Host string
	Port uint16
}

func (i index) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write(i.indexPageBytes)
	if err != nil {
		log.Println(err)
	}
}

func newIndexPage(port uint16) *index {
	t, err := template.ParseFiles("static/index.gohtml")
	if err != nil {
		log.Panic(err)
	}
	buff := &bytes.Buffer{}
	hostname, err := os.Hostname()
	if err != nil {
		log.Panic(err)
	}

	params := &indexPageParams{
		Host: hostname,
		Port: port,
	}

	err = t.Execute(buff, params)
	if err != nil {
		log.Panic(err)
	}

	return &index{indexPageBytes: buff.Bytes()}
}
