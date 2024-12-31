package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MedicinePostgres struct {
	db *sqlx.DB
}

func NewMedicinePostgres(db *sqlx.DB) *MedicinePostgres {
	return &MedicinePostgres{db: db}
}

func (r *MedicinePostgres) Create(medicine structures.Medicine) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description) values ($1, $2) RETURNING id", medicinesTable)
	row := r.db.QueryRow(query, medicine.Name, medicine.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MedicinePostgres) GetAll() ([]structures.Medicine, error) {
	var medicines []structures.Medicine
	query := fmt.Sprintf("SELECT * FROM %s", medicinesTable)
	err := r.db.Select(&medicines, query)
	if err != nil {
		return nil, err
	}
	return medicines, nil
}

func (r *MedicinePostgres) Get(id int) (structures.Medicine, error) {
	var medicine structures.Medicine
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", medicinesTable)
	err := r.db.Get(&medicine, query, id)
	if err != nil {
		return medicine, err
	}
	return medicine, nil
}

func (r *MedicinePostgres) Update(id int, input structures.UpdateMedicineInput) error {
	query := "UPDATE medicines SET"
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

func (r *MedicinePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", medicinesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
