package main

import "log"

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func checkErr(err error) {
	if err != nil {
		log.Printf("FAILURE MAIN: %v", err)
	}
}