package repository

import(
	"encoding/json"
	"github.com/github-timall/stat/entity"
)

func RedirectCreate(r *entity.Redirect) {
	r.ParseSub()

	SubCreate(&r.Sub1)
	SubCreate(&r.Sub2)
	SubCreate(&r.Sub3)
	SubCreate(&r.Sub4)
	SubCreate(&r.Sub5)

	SubUserCreate(r)

	var lastInsertId int
	err = db.QueryRow(`INSERT INTO redirect_fact(
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
		sub5_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id;`,
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
		r.Sub5.Id).Scan(&lastInsertId)
	checkErr(err)
	r.Id = lastInsertId
}

func RedirectFind(uuid string) entity.RedirectFact {
	var jsonPayload string

	err = db.QueryRow("SELECT row_to_json(redirect_fact) FROM redirect_fact WHERE uuid = $1 LIMIT 1;", uuid).Scan(&jsonPayload)
	checkErr(err)

	var r entity.RedirectFact

	err = json.Unmarshal([]byte(jsonPayload), &r)
	checkErr(err)

	return r
}

func SubCreate(s *entity.Sub) {
	if len(s.Name) > 0 {
		var lastInsertId int
		err = db.QueryRow(`SELECT id FROM sub WHERE name = $1 AND type = $2`, s.Name, s.Type).Scan(&lastInsertId)
		checkErr(err)
		if lastInsertId == 0 {
			err = db.QueryRow(`INSERT INTO sub (name, type) VALUES($1, $2) RETURNING id;`, s.Name, s.Type).Scan(&lastInsertId)
			checkErr(err)
		}
		s.Id = lastInsertId
	}
}

func SubUserCreate(r *entity.Redirect) {
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
