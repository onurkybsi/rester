package model

import "io"

// ReqModel struct
type ReqModel struct {
	TargetServerURL string    `json:"targetServerUrl"`
	Method          string    `json:"method"`
	ReqBody         io.Reader `json:"reqBody"`
}
