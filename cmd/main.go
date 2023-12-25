package main

import (
	"fmt"
	"log"
	"net/http"

	image "github.com/dadil/mosaicgenerator/httpserver"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			image.ServeForm(w)
		} else if r.Method == "POST" {
			image.HandlePostRequest(w, r)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
