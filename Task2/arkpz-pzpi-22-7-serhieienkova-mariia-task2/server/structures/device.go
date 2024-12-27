package structures

type Device struct {
	Id           int    `json:"id" db:"id"`
	Password     string `json:"password" binding:"required"`
	PasswordHash string `json:"-" db:"password_hash"`
}
