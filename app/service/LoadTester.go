package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/onurkybsi/rester/app/model"
)

var tr = &http.Transport{
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

type sendReqCallback func(simultaneousReqModel *model.SimultaneousReqModel)

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

func createReq(reqModel *model.ReqModel) (*http.Request, error) {
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

func sendReq(request *http.Request) (*model.ResModel, error) {
	result := &model.ResModel{}

	start := time.Now()
	res, err := client.Do(request)
	elapsed := int64(time.Since(start) / time.Millisecond)

	result.TimeSpent = elapsed

	if err == nil {
		result.DidErrOccured = false
		result.Status = res.Status
	} else {
		result.TimeSpent = 0
		result.DidErrOccured = true
		result.ErrMessage = err.Error()
	}

	return result, err
}

// SendSequentialReq send sequential requests
func SendSequentialReq(sequentialReqModel model.SequentialReqModel) model.TestResult {
	result := model.TestResult{
		IsOperationSuccess: true,
		Responses:          []model.ResModel{}}

	var totalElapsedTime int64
	var timeSpan time.Duration = time.Duration(sequentialReqModel.TimeSpanAsMs)
	var numberWithoutErrors int64

	for i := 0; i < sequentialReqModel.NumberOfReq; i++ {
		req, err := createReq(&sequentialReqModel.ReqModel)

		if err != nil {
			result.Responses = append(result.Responses, model.ResModel{TimeSpent: 0, DidErrOccured: true, ErrMessage: err.Error()})
			continue
		}

		res, err := sendReq(req)

		if err == nil {
			numberWithoutErrors++
			totalElapsedTime += res.TimeSpent

			result.Responses = append(result.Responses, *res)
		} else {
			result.Responses = append(result.Responses, *res)
		}

		if sequentialReqModel.TimeSpanAsMs > 0 {
			time.Sleep(timeSpan)
		}
	}

	result.AvgElapsedMs = totalElapsedTime / numberWithoutErrors

	return result
}

// SendMultipleReqSimultaneously send multiple requests simultaaneously
func SendMultipleReqSimultaneously(simultaneousReqModel model.SimultaneousReqModel) {
	result := model.TestResult{
		IsOperationSuccess: true,
		Responses:          []model.ResModel{}}

	var wg sync.WaitGroup

	context := model.SimultaneousReqContext{ReqModel: &simultaneousReqModel.ReqModel, Responses: result.Responses, WaitGroup: &wg}
	fmt.Println(context)

	for i := 0; i < simultaneousReqModel.NumberOfReq; i++ {
		wg.Add(1)

	}
}
