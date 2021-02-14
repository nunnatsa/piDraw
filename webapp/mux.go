package webapp

import (
	"net/http"
)

func GetMux(port uint16) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", newIndexPage(port))

	return mux
}
