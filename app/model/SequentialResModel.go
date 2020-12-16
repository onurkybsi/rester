package model

// SequentialResModel struct
type SequentialResModel struct {
	IsOperationSuccess bool       `json:"isOperationSuccess"`
	Responses          []ResModel `json:"resModel"`
	AvgElapsedMs       int64      `json:"avgElapsedMs"`
}
