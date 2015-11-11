package main

import "strconv"

// map of node ids -> nodes
var nodes map[string]Node
var nextNodeId int = 1

// map of node ids -> evident items 
var evidence map[string]EvidenceItems
var nextEvidenceId int = 1

func InitData() {
	nodes = make(map[string]Node)
	evidence = make(map[string]EvidenceItems)
	LoadData()
}

func AddNode(n Node) (Node) {
	n.Id = strconv.Itoa(nextNodeId)
	nodes[n.Id] = n
	nextNodeId++
	SaveData()
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

func AddEvidence(e EvidenceItem) {
	e.Id = strconv.Itoa(nextEvidenceId)
	evidence[e.NodeId] = append(evidence[e.NodeId], e)
	nextEvidenceId++
	SaveData()
}

func GetEvidence(nodeId string) (EvidenceItems) {
  return evidence[nodeId]
}
