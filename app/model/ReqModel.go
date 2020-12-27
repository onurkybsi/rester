package model

// ReqModel struct
type ReqModel struct {
	TargetServerURL string                 `json:"targetServerUrl"`
	Method          string                 `json:"method"`
	ReqBody         map[string]interface{} `json:"reqBody"`
	BearerToken     string                 `json:"bearerToken"`
}
