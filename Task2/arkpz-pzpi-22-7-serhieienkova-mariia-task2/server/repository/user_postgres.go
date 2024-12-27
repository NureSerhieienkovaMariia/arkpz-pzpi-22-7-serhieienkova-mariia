package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserActionPostgres struct {
	db *sqlx.DB
}

func NewUserActionPostgres(db *sqlx.DB) *UserActionPostgres {
	return &UserActionPostgres{db: db}
}

func (r *UserActionPostgres) CreateUser(user structures.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, surname, password_hash, premium_expiration_date) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	if user.PremiumExpirationDate == "" {
		user.PremiumExpirationDate = "0001-01-01T00:00:00Z"
	}
	row := r.db.QueryRow(query, user.Email, user.Name, user.Surname, user.PasswordHash, user.PremiumExpirationDate)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("error occured during 'create user' query: %w", err)
	}

	return id, nil
}

func (r *UserActionPostgres) GetAll() ([]structures.User, error) {
	var users []structures.User

	query := fmt.Sprintf(`SELECT * FROM %s`,
		usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		return users, fmt.Errorf("error occured during 'get all users' query: %w", err)
	}
	return users, err
}

func (r *UserActionPostgres) GetById(userId int) (structures.User, error) {
	var user structures.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`,
		usersTable)
	err := r.db.Get(&user, query, userId)
	if err != nil {
		return user, fmt.Errorf("error occured during 'get user by id' query: %w", err)
	}
	return user, err
}

func (r *UserActionPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		usersTable)
	_, err := r.db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("error occured during 'delete user by id' query: %w", err)
	}
	return err
}
func (r *UserActionPostgres) Update(userId int, input structures.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, input.Surname)
		argId++
	}

	if input.Email != "" {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, input.Email)
		argId++
	}

	if input.PasswordHash != "" {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, input.PasswordHash)
		argId++
	}

	if input.PremiumExpirationDate != "" {
		setValues = append(setValues, fmt.Sprintf("premium_expiration_date=$%d", argId))
		args = append(args, input.PremiumExpirationDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",
		usersTable, setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error occured during 'update user' query: %w", err)
	}
	return err
}

func (r *UserActionPostgres) GetByEmail(email string) (structures.User, error) {
	var user structures.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE email = $1`, usersTable)
	err := r.db.Get(&user, query, email)
	if err != nil {
		return user, fmt.Errorf("error occured during 'get user by email' query: %w", err)
	}
	return user, err
}
