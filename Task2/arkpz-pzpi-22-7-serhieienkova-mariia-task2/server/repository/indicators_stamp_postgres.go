package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type IndicatorsStampPostgres struct {
	db *sqlx.DB
}

func NewIndicatorsStampPostgres(db *sqlx.DB) *IndicatorsStampPostgres {
	return &IndicatorsStampPostgres{db: db}
}

func (r *IndicatorsStampPostgres) Create(indicatorsStamp structures.IndicatorsStamp) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (device_id, patient_id, timestamp, pulse, systolic_blood_pressure, diastolic_blood_pressure, temperature) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", indicatorsStampsTable)
	row := r.db.QueryRow(query, indicatorsStamp.DeviceId, indicatorsStamp.PatientId, indicatorsStamp.Timestamp, indicatorsStamp.Pulse, indicatorsStamp.SystolicBloodPressure, indicatorsStamp.DiastolicBloodPressure, indicatorsStamp.Temperature)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *IndicatorsStampPostgres) GetAll() ([]structures.IndicatorsStamp, error) {
	var indicatorsStamps []structures.IndicatorsStamp
	query := fmt.Sprintf("SELECT * FROM %s", indicatorsStampsTable)
	err := r.db.Select(&indicatorsStamps, query)
	if err != nil {
		return nil, err
	}
	return indicatorsStamps, nil
}

func (r *IndicatorsStampPostgres) Get(id int) (structures.IndicatorsStamp, error) {
	var indicatorsStamp structures.IndicatorsStamp
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", indicatorsStampsTable)
	err := r.db.Get(&indicatorsStamp, query, id)
	if err != nil {
		return indicatorsStamp, err
	}
	return indicatorsStamp, nil
}

func (r *IndicatorsStampPostgres) Update(id int, input structures.IndicatorsStamp) error {
	query := fmt.Sprintf("UPDATE %s SET device_id=$1, patient_id=$2, timestamp=$3, pulse=$4, systolic_blood_pressure=$5, diastolic_blood_pressure=$6, temperature=$7 WHERE id=$8", indicatorsStampsTable)
	_, err := r.db.Exec(query, input.DeviceId, input.PatientId, input.Timestamp, input.Pulse, input.SystolicBloodPressure, input.DiastolicBloodPressure, input.Temperature, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *IndicatorsStampPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", indicatorsStampsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
