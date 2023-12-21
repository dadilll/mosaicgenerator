package main

import (
	"fmt"
	"log"
	"net/http"

	image "github.com/dadil/mosaicgenerator/httpserver"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request:", r.URL.Path)
		if r.Method == "GET" {
			image.ServeForm(w)
		} else if r.Method == "POST" {
			image.HandlePostRequest(w, r)
		}
	})

	fmt.Println("Server is running at https://localhost:8080")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
