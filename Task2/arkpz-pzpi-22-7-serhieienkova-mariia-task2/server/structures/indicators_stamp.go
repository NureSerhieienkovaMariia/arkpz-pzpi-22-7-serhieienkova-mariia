package structures

type IndicatorsStamp struct {
	Id                     int     `json:"id" db:"id"`
	DeviceId               int     `json:"device_id" binding:"required"`
	PatientId              int     `json:"patient_id" binding:"required"`
	Timestamp              string  `json:"timestamp" binding:"required"`
	Pulse                  int     `json:"pulse" binding:"required"`
	SystolicBloodPressure  int     `json:"systolic_blood_pressure" binding:"required"`
	DiastolicBloodPressure int     `json:"diastolic_blood_pressure" binding:"required"`
	Temperature            float64 `json:"temperature" binding:"required"`
}
