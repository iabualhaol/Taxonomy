package main

import "log"
import "net/http"
import "github.com/gorilla/mux"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/nodes", NodeIndexHandler)
	router.HandleFunc("/nodes/{id}", NodeViewHandler)
	router.HandleFunc("/edges", EdgeIndexHandler)
	router.HandleFunc("/edges/{id}", EdgeViewHandler)
	router.HandleFunc("/nodes/{id}/evidences", EvidenceIndexHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

