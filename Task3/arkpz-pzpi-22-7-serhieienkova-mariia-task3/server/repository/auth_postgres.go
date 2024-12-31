package repository

import (
	"fmt"

	"clinic/server/structures"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user structures.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, surname, password_hash, premium_expiration_date) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Name, user.Surname, user.PasswordHash, user.PremiumExpirationDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username string, password string) (structures.User, error) {
	var user structures.User
	fmt.Println(fmt.Sprintf("query username: %v, password: %v", username, password))
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	fmt.Println(fmt.Sprintf("query: %v", query))
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) GetUserById(userId int) (structures.User, error) {
	var user structures.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`,
		usersTable)
	err := r.db.Get(&user, query, userId)
	if err != nil {
		return user, fmt.Errorf("error occured during 'get user by id' query: %w", err)
	}
	return user, err
}
