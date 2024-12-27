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

func (r *DiagnosisPostgres) Update(id int, input structures.Diagnosis) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, description=$2 WHERE id=$3", diagnosesTable)
	_, err := r.db.Exec(query, input.Name, input.Description, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *DiagnosisPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", diagnosesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
