package entity

import "log"

func checkErr(err error) {
	if err != nil {
		log.Printf("FAILURE ENTITY: %v", err)
	}
}
