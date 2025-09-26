package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello!"))
	})

	log.Print("Starting server on :4444")
	if err := http.ListenAndServe(":4444", mux); err != nil {
		log.Fatal(err)
	}
}
