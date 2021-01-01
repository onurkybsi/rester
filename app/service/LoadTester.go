package service

import (
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

func processSimultaneousReq(context *model.SimultaneousReqContext) {
	req, err := CreateReq(context.ReqModel)

	if err != nil {
		context.Responses = append(context.Responses, model.ResModel{DidErrOccured: true, ErrMessage: err.Error()})

		return
	}

	res, err := SendReq(req)

	context.Responses = append(context.Responses, *res)

	defer context.WaitGroup.Done()
}

// SendSequentialReq send sequential requests
func SendSequentialReq(sequentialReqModel model.SequentialReqModel) model.TestResult {
	result := model.TestResult{
		IsOperationSuccess: true,
		Responses:          make([]model.ResModel, sequentialReqModel.NumberOfReq)}

	var totalElapsedTime int64
	var timeSpan time.Duration = time.Duration(sequentialReqModel.TimeSpanAsMs)
	var numberWithoutErrors int64

	for i := 0; i < sequentialReqModel.NumberOfReq; i++ {
		req, err := CreateReq(&sequentialReqModel.ReqModel)

		if err != nil {
			result.Responses = append(result.Responses, model.ResModel{TimeSpent: 0, DidErrOccured: true, ErrMessage: err.Error()})
			continue
		}

		res, err := SendReq(req)

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
func SendMultipleReqSimultaneously(simultaneousReqModel model.SimultaneousReqModel) model.TestResult {
	result := model.TestResult{
		IsOperationSuccess: true,
		Responses:          make([]model.ResModel, simultaneousReqModel.NumberOfReq)}

	var wg sync.WaitGroup

	context := model.SimultaneousReqContext{ReqModel: &simultaneousReqModel.ReqModel, Responses: result.Responses, WaitGroup: &wg}

	for i := 0; i < simultaneousReqModel.NumberOfReq; i++ {
		wg.Add(1)
		go processSimultaneousReq(&context)
	}
	wg.Wait()

	return result
}
