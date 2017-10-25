package entity

type (
	Client struct {
		Id				int		`json:"id"`
		Uuid			string	`json:"uuid"`
		PhoneGeoCode	string	`json:"phone_geo_code"`
		Age				int		`json:"age"`
		Gender			string	`json:"gender"`
		Region			string	`json:"region"`
	}
)
