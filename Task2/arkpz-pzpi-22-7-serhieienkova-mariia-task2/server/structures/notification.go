package structures

type Notification struct {
	Id               int    `json:"id" db:"id"`
	IndicatorStampId int    `json:"indicator_stamp_id" binding:"required"`
	Message          string `json:"message" binding:"required"`
	Timestamp        string `json:"timestamp" binding:"required"`
}
