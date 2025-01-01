package structures

type Patient struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Birthday    string `json:"birthday" db:"birthday" binding:"required"`
	DiagnosisId int    `json:"diagnosis_id" db:"diagnosis_id" binding:"required"`
}

type CreatePatientInput struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Birthday    string `json:"birthday" binding:"required"`
	DiagnosisId int    `json:"diagnosis_id" binding:"required"`
	Relatives   []int  `json:"relatives,omitempty"`
	Doctors     []int  `json:"doctors,omitempty"`
}

type UpdatePatientInput struct {
	Name        *string `json:"name,omitempty"`
	Surname     *string `json:"surname,omitempty"`
	Birthday    *string `json:"birthday,omitempty"`
	DiagnosisId *int    `json:"diagnosis_id,omitempty"`
}

type PatientFullInfo struct {
	Patient       Patient
	Diagnoses     []Diagnosis
	Medicines     []Medicine
	Devices       []Device
	Indicators    []IndicatorsStamp
	Notifications []Notification
}
