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

func (r *NotificationPostgres) Update(id int, input structures.Notification) error {
	query := fmt.Sprintf("UPDATE %s SET indicator_stamp_id=$1, message=$2, timestamp=$3 WHERE id=$4", notificationsTable)
	_, err := r.db.Exec(query, input.IndicatorStampId, input.Message, input.Timestamp, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *NotificationPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", notificationsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
