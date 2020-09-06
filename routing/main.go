package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello friend!")
	})

	bookRouter := router.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Books there")
	})

	bookRouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		getVars := mux.Vars(r)
		title := getVars["title"]

		fmt.Fprintf(w, "Book name is: %v", title)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
