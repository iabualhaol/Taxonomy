package main

import "log"

func Log(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
