package main

import "time"

type (
	OrderFact struct {
		Id				int		`json:"id"`
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
		RedirectId		int

		Client 			Client	`json:"client"`
		Time 			OrderTime
	}

	OrderTime struct {
		OrderId		int
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

	Client struct {
		ClientId		int		`json:"id"`
		PhoneGeoCode	string	`json:"phone_geo_code"`
		Age				int		`json:"age"`
		Gender			string	`json:"gender"`
		Region			string	`json:"region"`
	}
)

func (f *OrderFact) Analysis() {
	var t OrderTime
	t.OrderId = f.Id
	t.DateAnalysis(f.CreatedAt)
	f.Time = t
}

func (t *OrderTime) DateAnalysis(date string) {
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
