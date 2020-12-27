package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onurkybsi/rester/app/model"
	"github.com/onurkybsi/rester/app/service"
)

// ReqSequential provide seq req
func ReqSequential(w http.ResponseWriter, r *http.Request) {
	var sequentialReqModel model.SequentialReqModel

	err := json.NewDecoder(r.Body).Decode(&sequentialReqModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	fmt.Println(sequentialReqModel.ReqModel.ReqBody)

	res := service.SendSequentialReq(sequentialReqModel)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
