package structures

type IndicatorsStamp struct {
	Id                     int     `json:"id" db:"id"`
	DeviceId               int     `json:"device_id" binding:"required" db:"device_id"`
	Timestamp              string  `json:"timestamp"`
	Pulse                  int     `json:"pulse" binding:"required"`
	SystolicBloodPressure  int     `json:"systolic_blood_pressure" binding:"required" db:"systolic_blood_pressure"`
	DiastolicBloodPressure int     `json:"distolic_blood_pressure" binding:"required" db:"distolic_blood_pressure"`
	Temperature            float64 `json:"temperature" binding:"required"`
}
