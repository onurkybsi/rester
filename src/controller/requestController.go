package controller

import (
	"encoding/json"
	"net/http"

	"github.com/onurkybsi/rester/src/model"
	"github.com/onurkybsi/rester/src/service"
)

// Ping : Ping the target server
func Ping(w http.ResponseWriter, r *http.Request) {
	var targetRequestModel model.RequestModel

	err := json.NewDecoder(r.Body).Decode(&targetRequestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	res, err := service.Ping(targetRequestModel.Domain)

	if err != nil && res == 200 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("OK !")
	}
}
