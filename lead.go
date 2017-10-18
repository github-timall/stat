package main

import "time"

type Lead struct {
	Id        	int			`json:"id"`
	UUID	  	string		`json:"uuid"`
	OfferId   	int			`json:"offer_id"`
	UserId	  	int			`json:"user_id"`
	CreatedAt	time.Time	`json:"created_at"`
}

type Leads []Lead