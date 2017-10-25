package handlers

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"github.com/github-timall/stat/entity"
	"github.com/github-timall/stat/repository"
)

func VeinOrderCreate(w http.ResponseWriter, r *http.Request) {
	//request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var order entity.OrderFact

	//validate
	if err = json.Unmarshal(body, &order); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	repository.OrderCreate(&order)

	//response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order.AsArray()); err != nil {
		panic(err)
	}
}

func VeinOrderStatus(w http.ResponseWriter, r *http.Request) {
	//request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var order entity.OrderFact

	//validate
	if err = json.Unmarshal(body, &order); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	repository.OrderStatus(&order)

	//response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order.AsArray()); err != nil {
		panic(err)
	}
}

func VeinOrderPayment(w http.ResponseWriter, r *http.Request) {
	//request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var order entity.OrderFact

	//validate
	if err = json.Unmarshal(body, &order); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	repository.OrderPayment(&order)

	//response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order.AsArray()); err != nil {
		panic(err)
	}
}