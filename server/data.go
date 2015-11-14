package main

import "fmt"
import "strconv"

// map of node ids -> nodes
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
	e.Arrows = "To"
	edges[e.From] = append(edges[e.From], e)
	fmt.Println("Add edge: ", e)
	nextEdgeId++
	SaveEdges()
	return e
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
