package repository

import (
	"github.com/github-timall/stat/entity"
)

func OrderCreate(o *entity.OrderFact) {
	o.Parse()
	r := RedirectFind(o.TrackingUuid)
	OrderTimeCreate(&o.Time)

	var lastInsertId int

	err = db.QueryRow(`INSERT INTO order_fact(
		uuid,
		type,
		payment_type,
		tracking_uuid,
		method,
		offer_id,
		user_id,
		client_id,
		redirect_id,
		time_id,
		status,
		status_partner,
		payment,
		sum) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING id;`,
		o.Uuid,
		o.Type,
		o.PaymentType,
		o.TrackingUuid,
		o.GetMethod(),
		o.OfferId,
		o.UserId,
		o.ClientId,
		r.Id,
		o.Time.Id,
		o.Status,
		o.StatusPartner,
		o.Payment,
		o.Sum).Scan(&lastInsertId)
	checkErr(err)
	o.Id = lastInsertId
}

func OrderTimeCreate(t *entity.OrderTime) {
	var lastInsertId int
	err = db.QueryRow(`INSERT INTO order_time(time,day_id,month_id,quarter_id,year_id)VALUES($1,$2,$3,$4,$5);`,
		t.Time,t.DayId,t.MonthId,t.QuarterId,t.YearId).Scan(&lastInsertId)
	checkErr(err)
	t.Id = lastInsertId
}

func OrderStatus(o *entity.OrderFact) {
	_, err = db.Exec(`UPDATE order_fact SET status=$1,status_partner=$2,sum=$3 WHERE uuid=$4;`,
		o.Status,o.StatusPartner,o.Sum,o.Uuid)
	checkErr(err)
}

func OrderPayment(o *entity.OrderFact) {
	_, err = db.Exec(`UPDATE order_fact SET payment=$1,sum=$2 WHERE uuid=$3;`,
		o.Payment,o.Sum,o.Uuid)
	checkErr(err)
}
