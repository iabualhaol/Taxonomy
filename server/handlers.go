package main

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
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
		Node { Id: "5", Label: "Mistakes" },
		Node { Id: "6", Label: "Errors" },
		Node { Id: "7", Label: "Omissions" },
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
		Edge { Id: "1", From: "1", To: "2", Arrows: "To" },
		Edge { Id: "2", From: "1", To: "3", Arrows: "To" },
		Edge { Id: "3", From: "1", To: "4", Arrows: "To" },
		Edge { Id: "4", From: "2", To: "5", Arrows: "To" },
		Edge { Id: "5", From: "2", To: "6", Arrows: "To" },
		Edge { Id: "6", From: "2", To: "7", Arrows: "To" },
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

func NodeEvidenceIndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if (id % 2 == 0) {
		evidence := EvidenceItems {
			EvidenceItem { Id: "1", Vote: 5, Author: "Michael", 
		  	Reason: "Links machine learning and code reuse" },
			EvidenceItem { Id: "2", Vote: 1, Author: "Ibrahim",
				Reason: "The concept is misleading" },
		}
		if err := json.NewEncoder(w).Encode(evidence); err != nil {
			panic(err)
		}
	} else {
		evidence := EvidenceItems {
			EvidenceItem { Id: "3", Vote: 3, Author: "Michael", 
		  	Reason: "Great article" },
			EvidenceItem { Id: "4", Vote: 1, Author: "Ibrahim",
				Reason: "The concept is also misleading" },
		}
		if err := json.NewEncoder(w).Encode(evidence); err != nil {
			panic(err)
		}
	}
}
