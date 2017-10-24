package handlers

import (
	"net/http"
	"encoding/json"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode("ok"); err != nil {
		panic(err)
	}
}
