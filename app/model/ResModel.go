package model

// ResModel struct
type ResModel struct {
	// ResponseData string `json:"responseData"`
	TimeSpent   int64  `json:"timeSpent"`
	DidErrOccur bool   `json:"didErrOccur"`
	ErrMessage  string `json:"errMessage"`
}
