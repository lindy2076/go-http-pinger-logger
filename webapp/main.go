package main

import (
	"fmt"
	"log"
	"net/http"
)

var sendOK bool = true

// simple handler that returns 200 and 500 alternatively
func statusHandler(w http.ResponseWriter, r *http.Request) {
	if sendOK {
		fmt.Fprintf(w, ":) Response code: %d", http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, ":( Response code: %d", http.StatusInternalServerError)
	}
	sendOK = !sendOK
}

func main() {
	http.HandleFunc("/status", statusHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
