package main

import "encoding/json"
import "io/ioutil"
import "log"
import "os"

func SaveData() {
	file, err := os.Create("data/evidence.dat")
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(evidence)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(data)
	file.Close()
}

func LoadData() {
	data, err := ioutil.ReadFile("data/evidence.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &evidence)
	if err != nil {
		log.Fatal(err)
	}
}
