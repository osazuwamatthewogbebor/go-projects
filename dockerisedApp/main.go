package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", strings.TrimPrefix(html.EscapeString(r.URL.Path), "/"))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	fmt.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
