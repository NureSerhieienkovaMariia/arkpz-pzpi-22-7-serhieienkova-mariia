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

func (r *PatientPostgres) Update(id int, input structures.UpdatePatientInput) error {
	query := fmt.Sprintf("UPDATE %s SET", patientsTable)
	args := []interface{}{}
	argId := 1

	if input.Name != nil {
		query += fmt.Sprintf(" name=$%d,", argId)
		args = append(args, *input.Name)
		argId++
	}
	if input.Surname != nil {
		query += fmt.Sprintf(" surname=$%d,", argId)
		args = append(args, *input.Surname)
		argId++
	}
	if input.Birthday != nil {
		query += fmt.Sprintf(" birthday=$%d,", argId)
		args = append(args, *input.Birthday)
		argId++
	}
	if input.DiagnosisId != nil {
		query += fmt.Sprintf(" diagnosis_id=$%d,", argId)
		args = append(args, *input.DiagnosisId)
		argId++
	}

	query = query[:len(query)-1] // Remove the trailing comma
	query += fmt.Sprintf(" WHERE id=$%d", argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PatientPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", patientsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PatientPostgres) GetDiagnosesByPatientID(patientID int) ([]structures.Diagnosis, error) {
	var diagnoses []structures.Diagnosis
	query := `SELECT d.* FROM ` + diagnosesTable + ` d
              JOIN ` + patientsTable + ` p ON p.diagnosis_id = d.id
              WHERE p.id = $1`
	err := r.db.Select(&diagnoses, query, patientID)
	return diagnoses, err
}

func (r *PatientPostgres) GetMedicinesByPatientID(patientID int) ([]structures.Medicine, error) {
	var medicines []structures.Medicine
	query := `SELECT m.* FROM ` + medicinesTable + ` m
              JOIN ` + patientMedicineTable + ` pm ON pm.medicine_id = m.id
              WHERE pm.patient_id = $1`
	err := r.db.Select(&medicines, query, patientID)
	return medicines, err
}

func (r *PatientPostgres) GetDevicesByPatientID(patientID int) ([]structures.Device, error) {
	var devices []structures.Device
	query := `SELECT * FROM ` + devicesTable + ` WHERE patient_id = $1`
	err := r.db.Select(&devices, query, patientID)
	return devices, err
}

func (r *PatientPostgres) GetIndicatorsByPatientID(patientID int) ([]structures.IndicatorsStamp, error) {
	var indicators []structures.IndicatorsStamp
	query := `SELECT * FROM ` + indicatorsStampsTable + ` WHERE device_id IN 
              (SELECT id FROM ` + devicesTable + ` WHERE patient_id = $1)`
	err := r.db.Select(&indicators, query, patientID)
	return indicators, err
}

func (r *PatientPostgres) GetNotificationsByPatientID(patientID int) ([]structures.Notification, error) {
	var notifications []structures.Notification
	query := `SELECT n.* FROM ` + notificationsTable + ` n
              JOIN ` + indicatorsStampsTable + ` ind ON ind.id = n.indicator_stamp_id
              JOIN ` + devicesTable + ` d ON d.id = ind.device_id
              WHERE d.patient_id = $1`
	err := r.db.Select(&notifications, query, patientID)
	return notifications, err
}
