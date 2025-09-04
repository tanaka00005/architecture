package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"architecture/handler"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	handler.Handler()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
