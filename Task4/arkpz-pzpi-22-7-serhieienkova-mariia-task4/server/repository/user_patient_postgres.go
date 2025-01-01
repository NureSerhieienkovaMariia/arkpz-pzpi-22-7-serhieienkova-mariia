package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPatientPostgres struct {
	db *sqlx.DB
}

func NewUserPatientPostgres(db *sqlx.DB) *UserPatientPostgres {
	return &UserPatientPostgres{db: db}
}

func (r *UserPatientPostgres) Create(userPatient structures.UserPatient) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, patient_id) values ($1, $2) RETURNING id", userPatientsTable)
	row := r.db.QueryRow(query, userPatient.UserId, userPatient.PatientId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPatientPostgres) GetAll() ([]structures.UserPatient, error) {
	var userPatients []structures.UserPatient
	query := fmt.Sprintf("SELECT * FROM %s", userPatientsTable)
	err := r.db.Select(&userPatients, query)
	if err != nil {
		return nil, err
	}
	return userPatients, nil
}

func (r *UserPatientPostgres) Get(id int) (structures.UserPatient, error) {
	var userPatient structures.UserPatient
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userPatientsTable)
	err := r.db.Get(&userPatient, query, id)
	if err != nil {
		return userPatient, err
	}
	return userPatient, nil
}

func (r *UserPatientPostgres) Update(id int, input structures.UserPatient) error {
	query := fmt.Sprintf("UPDATE %s SET user_id=$1, patient_id=$2 WHERE id=$3", userPatientsTable)
	_, err := r.db.Exec(query, input.UserId, input.PatientId, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserPatientPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", userPatientsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
