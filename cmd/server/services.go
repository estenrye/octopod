package main

import (
	"github.com/estenrye/octopod/internal/stack"
	"log"
	"net/http"
)

func getServices(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload, err = stack.ListServicesJson()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte(payload))
		if err != nil {
			log.Println(err)
			log.Printf("Successfully wrote %d bytes", write)
		}
	}
}
