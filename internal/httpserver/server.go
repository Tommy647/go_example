package httpserver

import (
	"log"
	"net/http"
)

type HelloResponse struct {
}

// HandleHello as a http request
func HandleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello http request")
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		log.Println(err.Error())
	}
}
