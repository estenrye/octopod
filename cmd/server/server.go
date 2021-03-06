package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/summary", getSummary).Methods(http.MethodGet)
	r.HandleFunc("/summary/{name}", getSummary).Methods(http.MethodGet)
	r.HandleFunc("/services", getServices).Methods(http.MethodGet)
	r.HandleFunc("/services/{name}", getServices).Methods(http.MethodGet)
	log.Println("Laucning server at http://localhost:9042/services")
	log.Fatal(http.ListenAndServe(":9042", r))
}
