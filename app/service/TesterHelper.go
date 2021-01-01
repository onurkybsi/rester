package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/onurkybsi/rester/app/model"
)

// Ping ping the server for health
func Ping(targetServerURL string) (int, error) {
	req, err := http.NewRequest(model.HeadMethod, targetServerURL, nil)

	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	resp.Body.Close()

	return resp.StatusCode, nil
}

// CreateReq create request by ReqModel
func CreateReq(reqModel *model.ReqModel) (*http.Request, error) {
	requestByte, _ := json.Marshal(reqModel.ReqBody)
	req, err := http.NewRequest(reqModel.Method, reqModel.TargetServerURL, bytes.NewReader(requestByte))

	if err != nil {

		return nil, errors.New("Error occurred when request create")
	}

	var bearerToken string = "Bearer " + reqModel.BearerToken
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// SendReq send request by http.Request
func SendReq(request *http.Request) (*model.ResModel, error) {
	result := &model.ResModel{}

	start := time.Now()
	res, err := client.Do(request)

	if err != nil {
		result.DidErrOccured = true
		result.ErrMessage = err.Error()
	} else {
		result.TimeSpent = int64(time.Since(start) / time.Millisecond)
		result.DidErrOccured = false
		result.Status = res.Status
	}

	return result, err
}
