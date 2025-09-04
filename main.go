package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/tanaka00005/architecture/handler"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	handler.Login()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
