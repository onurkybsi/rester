package handler

import (
	"encoding/json"
	"net/http"

	"github.com/onurkybsi/rester/app/model"
	"github.com/onurkybsi/rester/app/service"
)

// Ping : Ping the target server
func Ping(w http.ResponseWriter, r *http.Request) {
	var requestModel model.RequestModel

	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	res, err := service.Ping(requestModel.TargetServerURL)

	if err != nil && res == 200 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("OK !")
	}
}
