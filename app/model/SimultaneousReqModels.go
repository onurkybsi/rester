package model

import "sync"

// SimultaneousReqModel struct
type SimultaneousReqModel struct {
	ReqModel    ReqModel `json:"reqModel"`
	NumberOfReq int      `json:"numberOfReq"`
}

// SimultaneousReqContext struct
type SimultaneousReqContext struct {
	ReqModel  *ReqModel
	Responses []ResModel
	WaitGroup *sync.WaitGroup
}
