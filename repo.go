package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	err error
)

const (
	DB_USER     = "stat"
	DB_PASSWORD = "stat"
	DB_NAME     = "stat"
)

func init() {
	dbInfo := fmt.Sprintf("host=127.0.0.1 port=5432 user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbInfo)
	checkErr(err)
}

//func RepoFindLeads() Leads {
//	var l Leads
//	rows, err := db.Query("SELECT id, offer_id, user_id, uuid, created_at FROM lead;")
//	checkErr(err)
//
//	var lead Lead
//
//	for rows.Next() {
//		err = rows.Scan(&lead.Id, &lead.OfferId, &lead.UserId, &lead.UUID, &lead.CreatedAt)
//		checkErr(err)
//		l = append(l, lead)
//	}
//	rows.Close()
//	return l
//}
//
//func RepoFindLead(id int) Lead {
//	var l Lead
//	return l
//}
//
//func RepoCreateLead(l Lead) Lead {
//	var lastInsertId int
//	err = db.QueryRow("INSERT INTO lead(offer_id,user_id,uuid) VALUES($1,$2,$3) returning id;", l.OfferId, l.UserId, l.UUID).Scan(&lastInsertId)
//	checkErr(err)
//	l.Id = lastInsertId
//	return l
//}

func RepoCreateTracking(t Tracking) Tracking {
	_, err = db.Exec("INSERT INTO tracking(uuid,unique_id) VALUES($1,$2);", t.Uuid, t.UniqueId)
	checkErr(err)
	return t
}
