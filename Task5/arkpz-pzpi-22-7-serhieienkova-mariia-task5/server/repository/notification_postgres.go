package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type NotificationPostgres struct {
	db *sqlx.DB
}

func NewNotificationPostgres(db *sqlx.DB) *NotificationPostgres {
	return &NotificationPostgres{db: db}
}

func (r *NotificationPostgres) Create(notification structures.Notification) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (indicator_stamp_id, message, timestamp) values ($1, $2, $3) RETURNING id", notificationsTable)
	row := r.db.QueryRow(query, notification.IndicatorStampId, notification.Message, notification.Timestamp)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *NotificationPostgres) GetAll() ([]structures.Notification, error) {
	var notifications []structures.Notification
	query := fmt.Sprintf("SELECT * FROM %s", notificationsTable)
	err := r.db.Select(&notifications, query)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationPostgres) Get(id int) (structures.Notification, error) {
	var notification structures.Notification
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", notificationsTable)
	err := r.db.Get(&notification, query, id)
	if err != nil {
		return notification, err
	}
	return notification, nil
}

func (r *NotificationPostgres) GetAllByPatientID(patientID int) ([]structures.Notification, error) {
	var notifications []structures.Notification
	query := fmt.Sprintf(`
        SELECT n.id, n.indicator_stamp_id, n.message, n.timestamp
        FROM %s n
        JOIN %s i ON n.indicator_stamp_id = i.id
        WHERE i.device_id IN (SELECT id FROM %s WHERE patient_id = $1)
    `, notificationsTable, indicatorsStampsTable, devicesTable)
	err := r.db.Select(&notifications, query, patientID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
