package structures

type Notification struct {
	Id               int    `json:"id" db:"id"`
	IndicatorStampId int    `json:"indicator_stamp_id" binding:"required" db:"indicator_stamp_id"`
	Message          string `json:"message" binding:"required"`
	Timestamp        string `json:"timestamp" binding:"required"`
}
