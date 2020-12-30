package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/onurkybsi/rester/app/model"
)

var tr = &http.Transport{
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

func ping(targetServerURL string) (int, error) {
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

func createSeqReq(sequentialReqModel *model.SequentialReqModel) (*http.Request, error) {
	requestByte, _ := json.Marshal(sequentialReqModel.ReqModel.ReqBody)
	req, err := http.NewRequest(sequentialReqModel.ReqModel.Method, sequentialReqModel.ReqModel.TargetServerURL, bytes.NewReader(requestByte))

	if err != nil {

		return nil, errors.New("Error occurred when request create")
	}

	var bearerToken string = "Bearer " + sequentialReqModel.ReqModel.BearerToken
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// SendSequentialReq send sequential requests
func SendSequentialReq(sequentialReqModel model.SequentialReqModel) model.SequentialResModel {
	result := model.SequentialResModel{
		IsOperationSuccess: true,
		Responses:          []model.ResModel{}}

	var totalElapsedTime int64
	var timeSpan time.Duration = time.Duration(sequentialReqModel.TimeSpanAsMs)
	var numberWithoutErrors int64

	for i := 0; i < sequentialReqModel.NumberOfReq; i++ {
		req, err := createSeqReq(&sequentialReqModel)

		if err != nil {
			result.Responses = append(result.Responses, model.ResModel{TimeSpent: 0, DidErrOccur: true, ErrMessage: err.Error()})
			continue
		}

		start := time.Now()
		res, err := client.Do(req)
		elapsed := int64(time.Since(start) / time.Millisecond)

		if err == nil {
			numberWithoutErrors++
			totalElapsedTime += elapsed

			result.Responses = append(result.Responses, model.ResModel{TimeSpent: elapsed, DidErrOccur: false, Status: res.Status})
		} else {
			result.Responses = append(result.Responses, model.ResModel{TimeSpent: elapsed, DidErrOccur: true, ErrMessage: err.Error()})
		}

		if sequentialReqModel.TimeSpanAsMs > 0 {
			time.Sleep(timeSpan)
		}
	}

	result.AvgElapsedMs = totalElapsedTime / numberWithoutErrors

	return result
}
