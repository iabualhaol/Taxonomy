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
	if err := json.NewEncoder(w).Encode(GetAllNodes()); err != nil {
		panic(err)
	}
}

func NodeViewHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	nodeId := vars["id"]
	if err := json.NewEncoder(w).Encode(GetNode(nodeId)); err != nil {
		panic(err)
	}
}

func CreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	label := r.FormValue("label")
	node := AddNode(Node { Label: label })
	if err := json.NewEncoder(w).Encode(node); err != nil {
		panic(err)
	}
}

func EdgeIndexHandler(w http.ResponseWriter, r *http.Request) {
	edgeList := Edges{}
	if err := json.NewEncoder(w).Encode(edgeList); err != nil {
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
	nodeId := vars["id"]
	items := GetEvidence(nodeId)
	fmt.Println("Evidence for node: ", nodeId, " is: ", items)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func NodeCreateEvidenceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeId := vars["id"]
	vote, _ := strconv.Atoi(r.FormValue("vote"))
	AddEvidence(EvidenceItem { NodeId: nodeId, Author: "Michael", Vote: vote, 
		Reason: r.FormValue("reason") })
	items := GetEvidence(nodeId)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}
