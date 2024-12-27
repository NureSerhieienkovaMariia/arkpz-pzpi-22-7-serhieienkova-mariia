package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PatientMedicinePostgres struct {
	db *sqlx.DB
}

func NewPatientMedicinePostgres(db *sqlx.DB) *PatientMedicinePostgres {
	return &PatientMedicinePostgres{db: db}
}

func (r *PatientMedicinePostgres) Create(patientMedicine structures.PatientMedicine) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (patient_id, medicine_id, date, schedule) values ($1, $2, $3, $4) RETURNING id", patientMedicineTable)
	row := r.db.QueryRow(query, patientMedicine.PatientId, patientMedicine.MedicineId, patientMedicine.Date, patientMedicine.Schedule)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PatientMedicinePostgres) GetAll() ([]structures.PatientMedicine, error) {
	var patientMedicines []structures.PatientMedicine
	query := fmt.Sprintf("SELECT * FROM %s", patientMedicineTable)
	err := r.db.Select(&patientMedicines, query)
	if err != nil {
		return nil, err
	}
	return patientMedicines, nil
}

func (r *PatientMedicinePostgres) Get(id int) (structures.PatientMedicine, error) {
	var patientMedicine structures.PatientMedicine
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", patientMedicineTable)
	err := r.db.Get(&patientMedicine, query, id)
	if err != nil {
		return patientMedicine, err
	}
	return patientMedicine, nil
}

func (r *PatientMedicinePostgres) Update(id int, input structures.PatientMedicine) error {
	query := fmt.Sprintf("UPDATE %s SET patient_id=$1, medicine_id=$2, date=$3, schedule=$4 WHERE id=$5", patientMedicineTable)
	_, err := r.db.Exec(query, input.PatientId, input.MedicineId, input.Date, input.Schedule, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PatientMedicinePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientMedicineTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
