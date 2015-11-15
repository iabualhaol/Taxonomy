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
	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(GetAllNodes()); err != nil {
		panic(err)
	}
}

func NodeViewHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	nodeId := vars["id"]
	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(GetNode(nodeId)); err != nil {
		panic(err)
	}
}

func CreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	label := r.FormValue("label")
	node := AddNode(Node { Label: label })
	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(node); err != nil {
		panic(err)
	}
}

func EdgeIndexHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(GetAllEdges()); err != nil {
		panic(err)
	}
}

func EdgeViewHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
	id := vars["id"]
	SetHeaders(w)
	fmt.Fprintln(w, "Edge:", id)
}

func CreateOrUpdateEdgeHandler(w http.ResponseWriter, r *http.Request) {
	var edge Edge
	id := r.FormValue("id")
	from := r.FormValue("from")
	to := r.FormValue("to")
	fmt.Println("Create or update edge:", id, "from:", from, "to:", to)
	if (id == "") {
		edge = AddEdge(Edge { From: from, To: to })
	} else {
		UpdateEdge(Edge { Id: id, From: from, To: to })
	}
	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(edge); err != nil {
		panic(err)
	}
}

func NodeEvidenceIndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeId := vars["id"]
	items := GetEvidence(nodeId)
	SetHeaders(w)
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
	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
