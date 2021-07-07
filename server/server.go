package server

import (
	"io"
	"log"
	"net/http"
	"time"
)

func Run() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", serverHandler)
	http.HandleFunc("/print", printHandler)
	log.Printf("HTTP Server running at 127.0.0.1%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func serverHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("Wohoo!! You just hit the server")
	log.Println("/ API hit")
	io.WriteString(w, "Hello Harsh! I'm the serverHandler\n")
}

func printHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("/print API hit")
	// log.Printf("Request body: %s", r.Body)
	io.WriteString(w, "Hello Harsh! I'm the printHandler\n")
}
