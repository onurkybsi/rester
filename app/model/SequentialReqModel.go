package model

// SequentialReqModel struct
type SequentialReqModel struct {
	ReqModel     ReqModel `json:"reqModel"`
	TimeSpanAsMs int      `json:"timeSpanAsMs"`
	NumberOfReq  int      `json:"numberOfReq"`
}
