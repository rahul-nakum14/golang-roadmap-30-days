package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/health", HealthHandler).Methods("GET")

	log.Println("Server starting on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home endpoint called")
	w.Write([]byte("Hello World"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is running"))
}
