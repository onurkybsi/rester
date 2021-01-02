package model

// SequentialReqModel struct
type SequentialReqModel struct {
	ReqModel     ReqModel `json:"reqModel"`
	TimeSpanAsMs int64    `json:"timeSpanAsMs"`
	NumberOfReq  int      `json:"numberOfReq"`
}

// SequentialReqContext struct
type SequentialReqContext struct {
	SequentialReqModel  *SequentialReqModel
	TotalElapsedTime    int64
	NumberWithoutErrors int64
}
