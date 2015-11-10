package main

import "log"
import "net/http"
import "github.com/gorilla/mux"

func InitRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/nodes", NodeIndexHandler)
	router.HandleFunc("/nodes/{id}", NodeViewHandler)
	router.HandleFunc("/edges", EdgeIndexHandler)
	router.HandleFunc("/edges/{id}", EdgeViewHandler)
	router.HandleFunc("/nodes/{id}/evidence", NodeCreateEvidenceHandler).
		Methods("POST")
	router.HandleFunc("/nodes/{id}/evidence", NodeEvidenceIndexHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	InitData()
  InitRouter()
}
