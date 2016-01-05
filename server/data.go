package main

import "fmt"
import "strconv"

// map of node ids -> nodes
// Example--> http://www.dotnetperls.com/map-go
var nodes map[string]Node
var nextNodeId int = 1

//map of node ids -> edge lists
var edges map[string]Edges
var nextEdgeId int = 1

// map of node ids -> evident items 
var evidence map[string]EvidenceItems
var nextEvidenceId int = 1

func InitData() {
	nodes = make(map[string]Node)
	edges = make(map[string]Edges)
	evidence = make(map[string]EvidenceItems)
	LoadData()
}

func AddNode(n Node) (Node) {
	n.Id = strconv.Itoa(nextNodeId)
	nodes[n.Id] = n
	fmt.Println("Add node:", n)
	nextNodeId++
	SaveNodes()
	return n
}

func GetNode(nodeId string) (Node) {
	return nodes[nodeId]
}

func GetAllNodes() (Nodes) {
  nodeList := Nodes{}
	for _, node := range nodes {
		nodeList = append(nodeList, node)
  }
	return nodeList
}

func AddEdge(e Edge) (Edge) {
	e.Id = strconv.Itoa(nextEdgeId)
	e.Arrows = "From"
	e.Label = "is-a"
	edges[e.From] = append(edges[e.From], e)
	fmt.Println("Add edge: ", e)
	nextEdgeId++
	SaveEdges()
	return e
}

func UpdateEdge(e Edge) {
	// remove edge from old node (edge.From)
	nodeId, index, _ := GetEdge(e.Id)
	edges[nodeId] = append(edges[nodeId][:index], edges[nodeId][index:]...)
	// add edge to new node (e.From)
	edges[e.From] = append(edges[e.From], e)
	SaveEdges()
}

func GetEdge(edgeId string) (string, int, Edge) {
	for nodeId, edgesFromNode := range edges {
		for edgeIndex, edge := range edgesFromNode {
			if (edge.Id == edgeId) {
				return nodeId, edgeIndex, edge
			}
		}
	}
	return "", 0, Edge{}
}

func GetAllEdges() (Edges) {
	edgeList := Edges{}
	for _, edgesFromNode := range edges {
		for _, edge := range edgesFromNode {
			edgeList = append(edgeList, edge)
		}
	}
	return edgeList
}

func AddEvidence(e EvidenceItem) {
	e.Id = strconv.Itoa(nextEvidenceId)
	evidence[e.NodeId] = append(evidence[e.NodeId], e)
	nextEvidenceId++
	SaveEvidence()
}

func GetEvidence(nodeId string) (EvidenceItems) {
  return evidence[nodeId]
}
