package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DevicePostgres struct {
	db *sqlx.DB
}

func NewDevicePostgres(db *sqlx.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}

func (r *DevicePostgres) Create(device structures.Device) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (password_hash) values ($1) RETURNING id", devicesTable)
	row := r.db.QueryRow(query, device.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DevicePostgres) GetAll() ([]structures.Device, error) {
	var devices []structures.Device
	query := fmt.Sprintf("SELECT * FROM %s", devicesTable)
	err := r.db.Select(&devices, query)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *DevicePostgres) Get(id int) (structures.Device, error) {
	var device structures.Device
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", devicesTable)
	err := r.db.Get(&device, query, id)
	if err != nil {
		return device, err
	}
	return device, nil
}

func (r *DevicePostgres) Update(id int, input structures.Device) error {
	query := fmt.Sprintf("UPDATE %s SET password_hash=$1 WHERE id=$2", devicesTable)
	_, err := r.db.Exec(query, input.PasswordHash, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *DevicePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", devicesTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
