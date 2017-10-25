package entity

type (
	Redirect struct {
		Id			int
		Uuid		string		`json:"uuid"`
		UniqueId	int 		`json:"unique_id"`
		OfferId		int			`json:"offer_id"`
		AffiliateId	int			`json:"affiliate_id"`
		Landing		Landing		`json:"landing"`
		Prelanding	Prelanding	`json:"prelanding"`
		Location	Location	`json:"location"`
		Stream		Stream		`json:"stream"`
		Device		Device		`json:"device"`
		SubName1	string		`json:"sub1"`
		SubName2	string		`json:"sub2"`
		SubName3	string		`json:"sub3"`
		SubName4	string		`json:"sub4"`
		SubName5	string		`json:"sub5"`
		Sub1		Sub
		Sub2		Sub
		Sub3		Sub
		Sub4		Sub
		Sub5		Sub
	}

	RedirectFact struct {
		Id 				int
		Uuid 			string
		UniqueId 		int
		LandingId 		int
		PrelandingId	int
		StreamId 		int
		IsMobile 		bool
		GeoCode 		string
	}

	Landing struct {
		Id	int		`json:"id"`
		Url	string	`json:"url"`
	}

	Prelanding struct {
		Id	int		`json:"id"`
		Url	string	`json:"url"`
	}

	Location struct {
		Id		int		`json:"id"`
		Ip		string	`json:"ip"`
		GeoCode	string	`json:"geo_code"`
		City	string	`json:"city"`
	}

	Stream struct {
		Id		int		`json:"id"`
		Code	string	`json:"code"`
	}

	Device struct {
		Id			int		`json:"id"`
		IsMobile 	bool	`json:"is_mobile"`
	}

	Sub struct {
		Id 		int		`json:"id"`
		Name 	string	`json:"name"`
		Type 	int		`json:"type"`
	}

	SubUser struct {
		SubId 	int	`json:"sub_id"`
		UserId 	int	`json:"user_id"`
	}
)

func (r *Redirect) ParseSub() {
	r.Sub1.Name = r.SubName1
	r.Sub1.Type = 1

	r.Sub2.Name = r.SubName2
	r.Sub2.Type = 2

	r.Sub3.Name = r.SubName3
	r.Sub3.Type = 3

	r.Sub4.Name = r.SubName4
	r.Sub4.Type = 4

	r.Sub5.Name = r.SubName5
	r.Sub5.Type = 5
}

func (r Redirect) AsArray() map[string]interface{} {
	a := map[string]interface{}{"id": r.Id}
	return a
}
