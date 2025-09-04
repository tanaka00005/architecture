package handler

import (
	"fmt"
	"html"

	"log"
	"net/http"
)

func Handler() {
	http.HandleFunc("/ooo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ooo, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
