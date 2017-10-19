package main

import (
	"github.com/asaskevich/govalidator"
)

type (
	Tracking struct {
		Uuid		string		`valid:"uuid" json:"uuid"`
		UniqueId	int 		`valid:"int" json:"unique_id"`
		OfferId		int			`valid:"int" json:"offer_id"`
		AffiliateId	int			`valid:"int" json:"affiliate_id"`
		Landing		Landing		`valid:"int" json:"landing"`
		Prelanding	Prelanding	`valid:"int" json:"prelanding"`
		Location	Location	`valid:"int" json:"location"`
		Stream		Stream		`valid:"int" json:"stream"`
		Device		Device		`valid:"int" json:"device"`
		Sub1		string		`valid:"string" json:"sub1"`
		Sub2		string		`valid:"string" json:"sub2"`
		Sub3		string		`valid:"string" json:"sub3"`
		Sub4		string		`valid:"string" json:"sub4"`
		Sub5		string		`valid:"string" json:"sub5"`
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
)

func (t Tracking) Validate() (err error) {
	_, err = govalidator.ValidateStruct(t)
	_, err = govalidator.ValidateStruct(t.Landing)
	_, err = govalidator.ValidateStruct(t.Prelanding)
	_, err = govalidator.ValidateStruct(t.Location)
	_, err = govalidator.ValidateStruct(t.Stream)
	_, err = govalidator.ValidateStruct(t.Device)
	return
}
