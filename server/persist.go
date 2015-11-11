package main

import "encoding/json"
import "io/ioutil"
import "os"
import "strconv"

func SaveNodes() {
	file, err := os.Create("data/nodes.dat")
	Log(err)
	data, err := json.Marshal(nodes)
	Log(err)
	file.Write(data)
	file.Close()
}

func SaveEvidence() {
	file, err := os.Create("data/evidence.dat")
	Log(err)
	data, err := json.Marshal(evidence)
	Log(err)
	file.Write(data)
	file.Close()
}

func SaveEdges() {
	file, err := os.Create("data/edges.dat")
	Log(err)
	data, err := json.Marshal(edges)
	Log(err)
	file.Write(data)
	file.Close()
}

func LoadData() {
	LoadNodes()
	LoadEdges()
	LoadEvidence()
}

func LoadNodes() {
	data, err := ioutil.ReadFile("data/nodes.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &nodes)
	Log(err)
	UpdateNextNodeId()
}

func UpdateNextNodeId() {
	for _, n := range nodes {
		id, _ := strconv.Atoi(n.Id)
		if (nextNodeId <= id) {
			nextNodeId = id + 1
		}
	}
}

func LoadEdges() {
	data, err := ioutil.ReadFile("data/edges.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &edges)
	Log(err)
	UpdateNextEdgeId()
}

func UpdateNextEdgeId() {
	for _, eFN := range edges {
		for _, e := range eFN {
			id, _ := strconv.Atoi(e.Id)
			if (nextEdgeId <= id) {
				nextEdgeId = id + 1
			}
		}
	}
}

func LoadEvidence() {
	data, err := ioutil.ReadFile("data/evidence.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &evidence)
	Log(err)
	UpdateNextEvidenceId()
}

func UpdateNextEvidenceId() {
	for _, items := range evidence {
		for _, e := range items {
			id, _ := strconv.Atoi(e.Id)
			if (nextEvidenceId <= id) {
				nextEvidenceId = id + 1
			}
		}
	}
}
