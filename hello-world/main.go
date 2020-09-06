package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
