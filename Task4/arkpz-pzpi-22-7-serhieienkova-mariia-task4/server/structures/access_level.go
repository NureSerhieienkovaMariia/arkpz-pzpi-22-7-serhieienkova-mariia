package structures

type AccessLevel struct {
	Id   int    `json:"id" db:"id" binding:"required"`
	Name string `json:"name"`
}
