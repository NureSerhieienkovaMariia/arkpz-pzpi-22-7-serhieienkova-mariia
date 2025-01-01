package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DiagnosisPostgres struct {
	db *sqlx.DB
}

func NewDiagnosisPostgres(db *sqlx.DB) *DiagnosisPostgres {
	return &DiagnosisPostgres{db: db}
}

func (r *DiagnosisPostgres) Create(diagnosis structures.Diagnosis) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description) values ($1, $2) RETURNING id", diagnosesTable)
	row := r.db.QueryRow(query, diagnosis.Name, diagnosis.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DiagnosisPostgres) GetAll() ([]structures.Diagnosis, error) {
	var diagnoses []structures.Diagnosis
	query := fmt.Sprintf("SELECT * FROM %s", diagnosesTable)
	err := r.db.Select(&diagnoses, query)
	if err != nil {
		return nil, err
	}
	return diagnoses, nil
}

func (r *DiagnosisPostgres) Get(id int) (structures.Diagnosis, error) {
	var diagnosis structures.Diagnosis
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", diagnosesTable)
	err := r.db.Get(&diagnosis, query, id)
	if err != nil {
		return diagnosis, err
	}
	return diagnosis, nil
}

func (r *DiagnosisPostgres) Update(id int, input structures.UpdateDiagnosisInput) error {
	query := "UPDATE diagnoses SET"
	args := []interface{}{}
	argId := 1

	if input.Name != nil {
		query += fmt.Sprintf(" name=$%d,", argId)
		args = append(args, *input.Name)
		argId++
	}
	if input.Description != nil {
		query += fmt.Sprintf(" description=$%d,", argId)
		args = append(args, *input.Description)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *DiagnosisPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", diagnosesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
