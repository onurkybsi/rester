package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/onurkybsi/rester/app/model"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

// Ping : Send ping for checking server status.
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

// SendSequentialReq send sequential requests
func SendSequentialReq(sequentialReqModel model.SequentialReqModel) model.SequentialResModel {
	result := model.SequentialResModel{
		IsOperationSuccess: true,
		Responses:          make([]model.ResModel, sequentialReqModel.NumberOfReq),
	}

	var totalElapsedTime int64

	requestByte, _ := json.Marshal(sequentialReqModel.ReqModel.ReqBody)
	req, err := http.NewRequest(sequentialReqModel.ReqModel.Method, sequentialReqModel.ReqModel.TargetServerURL, bytes.NewReader(requestByte))

	if err != nil {
		result.IsOperationSuccess = false

		return result
	}

	var bearerToken string = "Bearer " + sequentialReqModel.ReqModel.BearerToken
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json; charset=utf8")

	for i := 0; i < sequentialReqModel.NumberOfReq; i++ {
		start := time.Now()
		res, err := client.Do(req)
		elapsed := int64(time.Since(start) / time.Millisecond)

		fmt.Println(res)

		totalElapsedTime += elapsed

		result.Responses = append(result.Responses, model.ResModel{TimeSpent: elapsed, DidErrOccur: err != nil, ErrMessage: "err occured when request to target"})

		if sequentialReqModel.TimeSpanAsMs > 0 {
			time.Sleep(time.Duration(sequentialReqModel.TimeSpanAsMs))
		}
	}

	result.AvgElapsedMs = totalElapsedTime / int64(sequentialReqModel.NumberOfReq)

	return result
}
