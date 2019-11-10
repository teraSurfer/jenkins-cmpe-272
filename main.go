package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is running.")
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Status).Methods("GET")

	log.Fatal(http.ListenAndServe(":3030", r))
}

func main() {
	fmt.Printf("Application is running on port: %d", 3030)
	handleRequests()
}
