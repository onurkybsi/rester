package model

// TestResult struct
type TestResult struct {
	IsOperationSuccess bool       `json:"isOperationSuccess"`
	Responses          []ResModel `json:"resModel"`
	AvgElapsedMs       int64      `json:"avgElapsedMs"`
}
