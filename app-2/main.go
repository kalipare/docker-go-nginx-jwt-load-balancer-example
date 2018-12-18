package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8082"

	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from " + port))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":" + port

	log.Println("Starting server at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
