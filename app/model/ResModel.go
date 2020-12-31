package model

// ResModel struct
type ResModel struct {
	// ResponseData string `json:"responseData"`
	TimeSpent     int64  `json:"timeSpent"`
	DidErrOccured bool   `json:"didErrOccured"`
	ErrMessage    string `json:"errMessage"`
	Status        string `json:"status"`
}
