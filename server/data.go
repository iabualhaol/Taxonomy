package main

import "strconv"

// map of node ids -> evident items 
var evidence map[string]EvidenceItems
var nextEvidenceId int = 1

func InitData() {
	evidence = make(map[string]EvidenceItems)
	LoadData()
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
