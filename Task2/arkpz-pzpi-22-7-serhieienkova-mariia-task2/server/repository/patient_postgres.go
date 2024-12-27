package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PatientPostgres struct {
	db *sqlx.DB
}

func NewPatientPostgres(db *sqlx.DB) *PatientPostgres {
	return &PatientPostgres{db: db}
}

func (r *PatientPostgres) Create(patient structures.Patient) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, birthday, diagnosis_id) values ($1, $2, $3, $4) RETURNING id", patientsTable)
	row := r.db.QueryRow(query, patient.Name, patient.Surname, patient.Birthday, patient.DiagnosisId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PatientPostgres) GetAll() ([]structures.Patient, error) {
	var patients []structures.Patient
	query := fmt.Sprintf("SELECT * FROM %s", patientsTable)
	err := r.db.Select(&patients, query)
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *PatientPostgres) GetById(id int) (structures.Patient, error) {
	var patient structures.Patient
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientsTable)
	err := r.db.Get(&patient, query, id)
	if err != nil {
		return patient, err
	}
	return patient, nil
}

func (r *PatientPostgres) Update(id int, input structures.Patient) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2, birthday=$3, diagnosis_id=$4 WHERE id=$5", patientsTable)
	_, err := r.db.Exec(query, input.Name, input.Surname, input.Birthday, input.DiagnosisId, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PatientPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
