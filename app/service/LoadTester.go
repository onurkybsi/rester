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

// SendSequentialReq send sequential requests
func SendSequentialReq(sequentialReqModel model.SequentialReqModel) model.TestResult {
	result := model.TestResult{
		IsOperationSuccess: true,
		Responses:          make([]model.ResModel, 0, sequentialReqModel.NumberOfReq)}

	context := getSequentialReqContext(&sequentialReqModel)

	processSequentialReq(context, &result)

	result.AvgElapsedMs = context.TotalElapsedTime / context.NumberWithoutErrors

	return result
}

func getSequentialReqContext(sequentialReqModel *model.SequentialReqModel) *model.SequentialReqContext {
	context := &model.SequentialReqContext{
		SequentialReqModel:  sequentialReqModel,
		TotalElapsedTime:    0,
		NumberWithoutErrors: 0}

	return context
}

func processSequentialReq(context *model.SequentialReqContext, result *model.TestResult) {
	for i := 0; i < context.SequentialReqModel.NumberOfReq; i++ {
		req, err := CreateReq(&context.SequentialReqModel.ReqModel)
		if err != nil {
			result.Responses = append(result.Responses, model.ResModel{DidErrOccured: true, ErrMessage: err.Error()})
			continue
		}

		res, err := SendReq(req)
		if err == nil {
			context.NumberWithoutErrors++
			context.TotalElapsedTime += res.TimeSpent
		}
		result.Responses = append(result.Responses, *res)

		if context.SequentialReqModel.TimeSpanAsMs > 0 {
			time.Sleep(time.Duration(context.SequentialReqModel.TimeSpanAsMs))
		}
	}
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
