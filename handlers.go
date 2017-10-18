package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LeadIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(RepoFindLeads()); err != nil {
		panic(err)
	}
}

func TodoView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var leadId int
	var err error
	if leadId, err = strconv.Atoi(vars["leadId"]); err != nil {
		panic(err)
	}
	lead := RepoFindLead(leadId)
	if lead.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(lead); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

//func TodoCreate(w http.ResponseWriter, r *http.Request) {
//	var todo Todo
//	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//	if err != nil {
//		panic(err)
//	}
//	if err := r.Body.Close(); err != nil {
//		panic(err)
//	}
//	if err := json.Unmarshal(body, &todo); err != nil {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(422)
//		if err := json.NewEncoder(w).Encode(err); err != nil {
//			panic(err)
//		}
//	}
//
//	t := RepoCreateTodo(todo)
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusCreated)
//	if err := json.NewEncoder(w).Encode(t); err != nil {
//		panic(err)
//	}
//}