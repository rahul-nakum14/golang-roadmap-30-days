package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000",r))

}
func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("helo world"))
}