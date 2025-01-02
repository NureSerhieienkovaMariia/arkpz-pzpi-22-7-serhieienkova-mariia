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

func (r *IndicatorsStampPostgres) Create(input structures.IndicatorsStamp) (int, error) {
	var id int
	query := `INSERT INTO indicators_stamps (device_id, timestamp, pulse, systolic_blood_pressure, distolic_blood_pressure, temperature) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRow(query, input.DeviceId, input.Timestamp, input.Pulse, input.SystolicBloodPressure, input.DiastolicBloodPressure, input.Temperature).Scan(&id)
	return id, err
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

func (r *IndicatorsStampPostgres) GetById(id int) (structures.IndicatorsStamp, error) {
	var indicatorsStamp structures.IndicatorsStamp
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", indicatorsStampsTable)
	err := r.db.Get(&indicatorsStamp, query, id)
	if err != nil {
		return indicatorsStamp, err
	}
	return indicatorsStamp, nil
}
