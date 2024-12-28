package structures

import (
	"errors"
)

type User struct {
	Id                    int    `json:"id" db:"id"`
	Name                  string `json:"name" binding:"required"`
	Surname               string `json:"surname" binding:"required"`
	PasswordHash          string `json:"password" db:"password_hash" binding:"required"`
	Email                 string `json:"email" binding:"required"`
	PremiumExpirationDate string `json:"premium_expiration_date" db:"premium_expiration_date"`
	AccessLevelId         string `json:"access_level_id" db:"access_level_id"`
}

type UserInfo struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name" binding:"required"`
	Surname string `json:"surname" db:"surname" binding:"required"`
	Email   string `json:"email" db:"email" binding:"required"`
}

type UpdateUserInput struct {
	Name                  string `json:"name"`
	Surname               string `json:"surname"`
	PasswordHash          string `json:"password" db:"password_hash"`
	Email                 string `json:"email"`
	PremiumExpirationDate string `json:"premium_expiration_date" db:"premium_expiration_date"`
	AccessLevelId         string `json:"access_level_id" db:"access_level_id"`
}

type UserToken struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

func (i User) Validate() error {
	if i.Name == "" || i.Surname == "" || i.PasswordHash == "" ||
		i.Email == "" {
		return errors.New("all fields should be filled")
	}

	if len(i.Name) > 32 {
		return errors.New("name should be less or equal 32")
	}

	if len(i.Surname) > 32 {
		return errors.New("surname should be less or equal 32")
	}

	return nil
}
