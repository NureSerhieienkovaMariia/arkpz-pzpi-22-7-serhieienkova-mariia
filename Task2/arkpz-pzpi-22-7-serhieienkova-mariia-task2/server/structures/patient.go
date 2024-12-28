package structures

type Patient struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Birthday    string `json:"birthday" db:"birthday" binding:"required"`
	DiagnosisId int    `json:"diagnosis_id" db:"diagnosis_id"`
}
