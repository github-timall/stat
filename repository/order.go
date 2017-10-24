package repository

import (
	"github.com/github-timall/stat/entity"
)

func OrderCreate(o entity.OrderFact) entity.OrderFact {
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

	OrderTimeCreate(o.Time)
	ClientCreate(o.Client)

	return o
}

func OrderTimeCreate(t entity.OrderTime) {
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

func ClientCreate(c entity.Client) {
	_, err = db.Exec(`INSERT INTO client(client_id,phone_geo_code,age,gender,region) VALUES ($1,$2,$3,$4,$5);`,
		c.ClientId,c.PhoneGeoCode,c.Age,c.Gender,c.Region)
	checkErr(err)
}

func OrderStatus(o entity.OrderFact) entity.OrderFact {
	_, err = db.Exec(`UPDATE order_fact SET status=$1,status_partner=$2 WHERE order_id=$3;`,o.Status,o.StatusPartner,o.Id)
	checkErr(err)
	RepoClientUpdate(o.Client)
	return o
}

func OrderPayment(o entity.OrderFact) entity.OrderFact {
	_, err = db.Exec(`UPDATE order_fact SET payment=$1 WHERE order_id=$2;`,o.Payment,o.Id)
	checkErr(err)
	RepoClientUpdate(o.Client)
	return o
}

func RepoClientUpdate(c entity.Client) {
	_, err = db.Exec(`UPDATE client SET phone_geo_code=$1,age=$2,gender=$3,region=$4 WHERE client_id=$5;`,
		c.PhoneGeoCode,c.Age,c.Gender,c.Region,c.ClientId)
	checkErr(err)
}
