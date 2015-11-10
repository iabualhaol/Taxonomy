package main

import "strconv"

var evidence map[string]EvidenceItem
var nextEvidenceId int = 1

func InitData() {
	evidence = make(map[string]EvidenceItem)
}

func AddEvidence(e EvidenceItem) {
	e.Id = strconv.Itoa(nextEvidenceId)
	nextEvidenceId++
	evidence[e.NodeId] = e
}

func GetEvidence(nodeId string) (EvidenceItem) {
  return evidence[nodeId]
}
