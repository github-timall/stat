package handlers

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"github.com/github-timall/stat/entity"
	"github.com/github-timall/stat/repository"
)

func VeinRedirectCreate(w http.ResponseWriter, r *http.Request) {
	//request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var redirect entity.Redirect

	//validate
	if err = json.Unmarshal(body, &redirect); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	repository.RedirectCreate(&redirect)

	//response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(redirect.AsArray()); err != nil {
		panic(err)
	}
}
