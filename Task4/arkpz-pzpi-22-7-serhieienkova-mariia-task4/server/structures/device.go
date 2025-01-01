package structures

type Device struct {
	Id           int    `json:"id" db:"id"`
	Password     string `json:"password" binding:"required"`
	PatientId    int    `json:"patient_id" binding:"required" db:"patient_id"`
	PasswordHash string `json:"-" db:"password_hash"`
}

type UpdateDeviceInput struct {
	Password     *string `json:"password,omitempty"`
	PatientId    *int    `json:"patient_id,omitempty"`
	PasswordHash string  `json:"-" db:"password_hash"`
}
