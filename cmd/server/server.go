package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/services", getServices).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9042", r))
}
