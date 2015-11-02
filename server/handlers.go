package main

import "fmt"
import "net/http"
import "encoding/json"
import "github.com/gorilla/mux"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Taxi (c)")
}

func NodeIndexHandler(w http.ResponseWriter, r *http.Request) {
	nodes := Nodes {
		Node { Id: "1", Label: "Actions" },
		Node { Id: "2", Label: "Inadvertent" },
		Node { Id: "3", Label: "Deliberate" },
		Node { Id: "4", Label: "Inaction" },
	}
	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		panic(err)
	}
}

func NodeViewHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "Node:", id)
}

func EdgeIndexHandler(w http.ResponseWriter, r *http.Request) {
	edges := Edges {
		Edge { Id: "1", From: "1", To: "2" },
		Edge { Id: "2", From: "1", To: "3" },
		Edge { Id: "3", From: "1", To: "4" },
	}
	if err := json.NewEncoder(w).Encode(edges); err != nil {
		panic(err)
	}
}

func EdgeViewHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "Edge:", id)
}
