package structures

type UserPatient struct {
	Id        int `json:"id" db:"id"`
	UserId    int `json:"user_id" binding:"required"`
	PatientId int `json:"patient_id" binding:"required"`
}
