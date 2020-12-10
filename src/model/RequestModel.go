package model

// RequestModel struct
type RequestModel struct {
	TargetServerURL string `json:"targetServeUrl"`
	Method          string `json:"method"`
	RequestBody     string `json:"requestBody"`
}
