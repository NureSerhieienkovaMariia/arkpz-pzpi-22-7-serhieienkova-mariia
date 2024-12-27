package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable            = "public.\"users\""
	patientsTable         = "public.\"patients\""
	medicinesTable        = "public.\"medicines\""
	diagnosesTable        = "public.\"diagnoses\""
	patientMedicineTable  = "public.\"patients_medicines\""
	userPatientsTable     = "public.\"users_patients\""
	devicesTable          = "public.\"devices\""
	notificationsTable    = "public.\"notifications\""
	indicatorsStampsTable = "public.\"indicators_stamps\""
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
