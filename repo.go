package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
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

func RepoRedirectCreate(r Redirect) Redirect {
	r.ParseSub()

	RepoSubCreate(&r.Sub1)
	RepoSubCreate(&r.Sub2)
	RepoSubCreate(&r.Sub3)
	RepoSubCreate(&r.Sub4)
	RepoSubCreate(&r.Sub5)

	RepoSubUserCreate(r)

	_, err = db.Exec(`INSERT INTO redirect_fact(
		uuid,
		unique_id,
		landing_id,
		prelanding_id,
		stream_id,
		device_is_mobile,
		location_geo_code,
		sub1_id,
		sub2_id,
		sub3_id,
		sub4_id,
		sub5_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`,
		r.Uuid,
		r.UniqueId,
		r.Landing.Id,
		r.Prelanding.Id,
		r.Stream.Id,
		r.Device.IsMobile,
		r.Location.GeoCode,
		r.Sub1.Id,
		r.Sub2.Id,
		r.Sub3.Id,
		r.Sub4.Id,
		r.Sub5.Id)
	checkErr(err)

	return r
}

func RepoRedirectFind(uuid string) RedirectFact {
	var jsonPayload string

	err = db.QueryRow("SELECT row_to_json(redirect_fact) FROM redirect_fact WHERE uuid = $1;", uuid).Scan(&jsonPayload)
	checkErr(err)

	var r RedirectFact

	err = json.Unmarshal([]byte(jsonPayload), &r)
	checkErr(err)

	return r
}

func RepoViewCreate(v View) View {
	_, err = db.Exec("INSERT INTO view_fact(uuid,unique_id) VALUES($1,$2);", v.Uuid, v.UniqueId)
	checkErr(err)
	return v
}

func RepoOrderCreate(o OrderFact) OrderFact {
	o.Analysis()

	r := RepoRedirectFind(o.TrackingUuid)

	_, err = db.Exec(`INSERT INTO order_fact(
		order_id,
		type,
		payment_type,
		tracking_uuid,
		offer_id,
		user_id,
		client_id,
		status,
		status_partner,
		payment,
		sum,
		redirect_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`,
		o.Id,
		o.Type,
		o.PaymentType,
		o.TrackingUuid,
		o.OfferId,
		o.UserId,
		o.ClientId,
		o.Status,
		o.StatusPartner,
		o.Payment,
		o.Sum,
		r.Id)
	checkErr(err)

	RepoOrderTimeCreate(o.Time)
	RepoClientCreate(o.Client)

	return o
}

func RepoOrderTimeCreate(t OrderTime) {
	_, err = db.Exec(`INSERT INTO order_time(
		order_id,
		time,
		day_id,
		month_id,
		quarter_id,
		year_id) VALUES ($1,$2,$3,$4,$5,$6);`,
		t.OrderId,
		t.Time,
		t.DayId,
		t.MonthId,
		t.QuarterId,
		t.YearId)
	checkErr(err)
}

func RepoClientCreate(c Client) {
	_, err = db.Exec(`INSERT INTO client(client_id,phone_geo_code,age,gender,region) VALUES ($1,$2,$3,$4,$5);`,
		c.ClientId,c.PhoneGeoCode,c.Age,c.Gender,c.Region)
	checkErr(err)
}

func RepoOrderStatus(o OrderFact) OrderFact {
	_, err = db.Exec(`UPDATE order_fact SET status=$1,status_partner=$2 WHERE order_id=$3;`,o.Status,o.StatusPartner,o.Id)
	checkErr(err)
	RepoClientUpdate(o.Client)
	return o
}

func RepoOrderPayment(o OrderFact) OrderFact {
	_, err = db.Exec(`UPDATE order_fact SET payment=$1 WHERE order_id=$2;`,o.Payment,o.Id)
	checkErr(err)
	RepoClientUpdate(o.Client)
	return o
}

func RepoClientUpdate(c Client) {
	_, err = db.Exec(`UPDATE client SET phone_geo_code=$1,age=$2,gender=$3,region=$4 WHERE client_id=$5;`,
		c.PhoneGeoCode,c.Age,c.Gender,c.Region,c.ClientId)
	checkErr(err)
}

func RepoSubCreate(s *Sub) {
	if len(s.Name) > 0 {
		var lastInsertId int
		err = db.QueryRow(`SELECT id FROM sub WHERE name = $1 AND type = $2`, s.Name, s.Type).Scan(&lastInsertId)
		checkErr(err)
		if lastInsertId == 0 {
			query := `INSERT INTO sub (name, type) VALUES($1, $2) RETURNING id;`
			err = db.QueryRow(query, s.Name, s.Type).Scan(&lastInsertId)
			checkErr(err)
		}
		s.Id = lastInsertId
	}
}

func RepoSubUserCreate(r Redirect)  {
	query := `INSERT INTO sub_user (sub_id, user_id) SELECT $1, $2
		WHERE NOT EXISTS (SELECT 1 FROM sub_user WHERE sub_id = $3 AND user_id = $4);`

	if r.Sub1.Id > 0 {
		_, err = db.Exec(query, r.Sub1.Id, r.AffiliateId, r.Sub1.Id, r.AffiliateId)
		checkErr(err)
	}

	if r.Sub2.Id > 0 {
		_, err = db.Exec(query, r.Sub2.Id, r.AffiliateId, r.Sub2.Id, r.AffiliateId)
		checkErr(err)
	}

	if r.Sub3.Id > 0 {
		_, err = db.Exec(query, r.Sub3.Id, r.AffiliateId, r.Sub3.Id, r.AffiliateId)
		checkErr(err)
	}

	if r.Sub4.Id > 0 {
		_, err = db.Exec(query, r.Sub4.Id, r.AffiliateId, r.Sub4.Id, r.AffiliateId)
		checkErr(err)
	}

	if r.Sub5.Id > 0 {
		_, err = db.Exec(query, r.Sub5.Id, r.AffiliateId, r.Sub5.Id, r.AffiliateId)
		checkErr(err)
	}
}

func RepoSubFind(name string) Sub {
	var jsonPayload string

	var s Sub

	err = db.QueryRow("SELECT row_to_json(sub) FROM sub WHERE name = $1;", name).Scan(&jsonPayload)
	if err != nil {
		return s
	}

	err = json.Unmarshal([]byte(jsonPayload), &s)
	checkErr(err)

	return s
}