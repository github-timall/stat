package entity

import "time"

const (
	TYPE_JS = 1
	TYPE_API = 2
	TYPE_POSTBACK = 3
	TYPE_IFRAME = 4
)

type (
	OrderFact struct {
		Id				int
		Uid				string	`json:"uid"`
		Uuid			string	`json:"uuid"`
		Type			int		`json:"type"`
		PaymentType		int		`json:"payment_type"`
		TrackingUuid	string	`json:"tracking_uuid"`
		ClientId		int		`json:"client_id"`
		AddressId		int		`json:"address_id"`
		UserId			int		`json:"user_id"`
		Status			int		`json:"status"`
		StatusPartner	int		`json:"status_partner"`
		Payment			float64	`json:"payment"`
		Sum				float64	`json:"sum"`
		CreatedAt		string	`json:"created_at"`
		OfferId			int		`json:"offer_id"`

		Time 			OrderTime
	}

	OrderTime struct {
		Id		int
		Time		string
		DayId	 	int
		MonthId		time.Month
		QuarterId	int
		YearId		int
	}

	OrderCrm struct {
		OrderId		int		`json:"order_id"`
		ExternalId 	int		`json:"external_id"`
		CrmId		int		`json:"crm_id"`
	}
)

func (o OrderFact) GetMethod() int {
	if o.Type == TYPE_JS {
		return  2
	}

	if o.Type == TYPE_API {
		return  4
	}

	if o.Type == TYPE_POSTBACK {
		return  5
	}

	if o.Type == TYPE_IFRAME {
		return  3
	}
	return  2
}

func (o *OrderFact) Parse() {
	var t OrderTime
	t.ParseDate(o.CreatedAt)
	o.Time = t
}

func (t *OrderTime) ParseDate(date string) {
	layout := "2006-01-02 15:04:05"
	orderTime, err := time.Parse(layout, date)
	checkErr(err)

	t.Time = orderTime.Format(layout)
	t.DayId = orderTime.Day()
	t.MonthId = orderTime.Month()
	t.YearId = orderTime.Year()

	quarter := 4
	switch orderTime.Month() {
	case time.January, time.February, time.March:
		quarter = 1
	case time.April, time.May, time.June:
		quarter = 2
	case time.July, time.August, time.September:
		quarter = 3
	}
	t.QuarterId = quarter
}

func (o OrderFact) AsArray() map[string]interface{} {
	a := map[string]interface{}{"id": o.Id}
	return a
}
