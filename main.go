package main

import (
	"log"
	"net/http"
)
type (
	Visitor struct {
		Host       string
		UserAgent  string
		RequestUri string
		Headers    http.Header
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
