package repository

import "github.com/github-timall/stat/entity"

func ClientCreate(c entity.Client) {
	_, err = db.Exec(`INSERT INTO client(client_id,phone_geo_code,age,gender,region) VALUES ($1,$2,$3,$4,$5);`,
		c.Id,c.PhoneGeoCode,c.Age,c.Gender,c.Region)
	checkErr(err)
}

func ClientUpdate(c entity.Client) {
	_, err = db.Exec(`UPDATE client SET phone_geo_code=$1,age=$2,gender=$3,region=$4 WHERE client_id=$5;`,
		c.PhoneGeoCode,c.Age,c.Gender,c.Region,c.Id)
	checkErr(err)
}