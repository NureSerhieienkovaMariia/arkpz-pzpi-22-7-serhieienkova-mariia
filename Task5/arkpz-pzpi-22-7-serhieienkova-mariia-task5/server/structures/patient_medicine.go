package structures

type PatientMedicine struct {
	Id         int    `json:"id" db:"id"`
	PatientId  int    `json:"patient_id" binding:"required"`
	MedicineId int    `json:"medicine_id" binding:"required"`
	Date       string `json:"date" db:"date" binding:"required"`
	Schedule   string `json:"schedule" binding:"required"`
}
