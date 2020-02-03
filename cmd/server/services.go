package main

import (
	"fmt"
	"github.com/estenrye/octopod/internal/stack"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getServices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(fmt.Sprintf("Request: http:localhost:9042/services/%s", vars["name"]))
	w.Header().Set("Content-Type", "application/json")
	var data, dataError = stack.ListServicesByName(vars["name"])

	if dataError != nil {
		log.Println(dataError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload, conversionError = data.ToJSON()

	if conversionError != nil {
		log.Println(conversionError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write([]byte(payload))
	if err != nil {
		log.Println(err)
		log.Printf("Successfully wrote %d bytes", write)
	}
}
