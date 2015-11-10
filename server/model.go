package main

type EvidenceItem struct {
  Id string `json:"id"`
	NodeId string `json:"nodeid"`
	Vote int `json:"vote"`
	Author string `json:"author"`
	Reason string `json:"reason"`
}

type EvidenceItems []EvidenceItem

type Node struct {
	Id string `json:"id"`
	Label string `json:"label"`
}

type Nodes []Node

type Edge struct {
	Id string `json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	Arrows string `json:"arrows"`
	Label string `json:"label"`
}

type Edges []Edge
