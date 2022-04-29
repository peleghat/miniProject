package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO - Remove later
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	// hi
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
