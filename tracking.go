package main

type (
	Redirect struct {
		Uuid		string		`valid:"uuid" json:"uuid"`
		UniqueId	int 		`valid:"int" json:"unique_id"`
		OfferId		int			`valid:"int" json:"offer_id"`
		AffiliateId	int			`valid:"int" json:"affiliate_id"`
		Landing		Landing		`valid:"int" json:"landing"`
		Prelanding	Prelanding	`valid:"int" json:"prelanding"`
		Location	Location	`valid:"int" json:"location"`
		Stream		Stream		`valid:"int" json:"stream"`
		Device		Device		`valid:"int" json:"device"`
		SubName1	string		`json:"sub1"`
		SubName2	string		`json:"sub2"`
		SubName3	string		`json:"sub3"`
		SubName4	string		`json:"sub4"`
		SubName5	string		`json:"sub5"`
		Sub1		Sub			`json:"subs1"`
		Sub2		Sub			`json:"subs2"`
		Sub3		Sub			`json:"subs3"`
		Sub4		Sub			`json:"subs4"`
		Sub5		Sub			`json:"subs5"`
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

	View struct {
		Uuid		string		`valid:"uuid" json:"uuid"`
		UniqueId	int 		`valid:"int" json:"unique_id"`
		OfferId		int			`valid:"int" json:"offer_id"`
		AffiliateId	int			`valid:"int" json:"affiliate_id"`
		Landing		Landing		`valid:"int" json:"landing"`
		Prelanding	Prelanding	`valid:"int" json:"prelanding"`
		Location	Location	`valid:"int" json:"location"`
		Stream		Stream		`valid:"int" json:"stream"`
		Device		Device		`valid:"int" json:"device"`
	}

	Landing struct {
		Id	int		`valid:"int" json:"id"`
		Url	string	`valid:"string" json:"url"`
	}

	Prelanding struct {
		Id	int		`valid:"int" json:"id"`
		Url	string	`valid:"string" json:"url"`
	}

	Location struct {
		Id		int		`valid:"int" json:"id"`
		Ip		string	`valid:"string" json:"ip"`
		GeoCode	string	`valid:"string" json:"geo_code"`
		City	string	`valid:"string" json:"city"`
	}

	Stream struct {
		Id		int		`valid:"int" json:"id"`
		Code	string	`valid:"string" json:"code"`
	}

	Device struct {
		Id			int		`valid:"int" json:"id"`
		IsMobile 	bool	`valid:"bool" json:"is_mobile"`
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