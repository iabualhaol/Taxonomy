package main

type Node struct {
	Id string `json:"id"`
	Label string `json:"label"`
}

type Nodes []Node

type Edge struct {
	Id string `json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	Label string `json:"label"`
}

type Edges []Edge
