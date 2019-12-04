package main

import (
	"encoding/json"
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
		jsonVisitorData, err := json.Marshal(
			Visitor{
				r.Host,
				r.Header.Get("User-Agent"),
				r.URL.Path,
				r.Header,
			})
		if err != nil {
			panic(err)
		}
		_, err = w.Write(jsonVisitorData)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
