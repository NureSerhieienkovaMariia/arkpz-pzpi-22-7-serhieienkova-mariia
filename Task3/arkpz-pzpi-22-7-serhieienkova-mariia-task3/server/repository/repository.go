package repository

import (
	"clinic/server/structures"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepo interface {
	CreateUser(user structures.User) (int, error)
	GetUser(username, password string) (structures.User, error)
	GetUserById(userId int) (structures.User, error)
}

type UserRepo interface {
	CreateUser(user structures.User) (int, error)
	GetAll() ([]structures.User, error)
	GetById(userId int) (structures.User, error)
	Delete(userId int) error
	Update(userId int, input structures.UpdateUserInput) error
	GetByEmail(email string) (structures.User, error)
}

type PatientRepo interface {
	Create(patient structures.Patient) (int, error)
	GetAll() ([]structures.Patient, error)
	GetById(id int) (structures.Patient, error)
	Update(id int, input structures.UpdatePatientInput) error
	Delete(id int) error
	GetDiagnosesByPatientID(patientID int) ([]structures.Diagnosis, error)
	GetMedicinesByPatientID(patientID int) ([]structures.Medicine, error)
	GetDevicesByPatientID(patientID int) ([]structures.Device, error)
	GetIndicatorsByPatientID(patientID int) ([]structures.IndicatorsStamp, error)
	GetNotificationsByPatientID(patientID int) ([]structures.Notification, error)
}

type MedicineRepo interface {
	Create(medicine structures.Medicine) (int, error)
	GetAll() ([]structures.Medicine, error)
	Get(id int) (structures.Medicine, error)
	Update(id int, input structures.UpdateMedicineInput) error
	Delete(id int) error
}

type DiagnosisRepo interface {
	Create(diagnosis structures.Diagnosis) (int, error)
	GetAll() ([]structures.Diagnosis, error)
	Get(id int) (structures.Diagnosis, error)
	Update(id int, input structures.UpdateDiagnosisInput) error
	Delete(id int) error
}

type PatientMedicineRepo interface {
	Create(patientMedicine structures.PatientMedicine) (int, error)
	GetAll() ([]structures.PatientMedicine, error)
	Get(id int) (structures.PatientMedicine, error)
	Update(id int, input structures.PatientMedicine) error
	Delete(id int) error
	SetMedicineToPatient(patientMedicine structures.PatientMedicine) error
}

type UserPatientRepo interface {
	Create(userPatient structures.UserPatient) (int, error)
	GetAll() ([]structures.UserPatient, error)
	Get(id int) (structures.UserPatient, error)
	Update(id int, input structures.UserPatient) error
	Delete(id int) error
}

type DeviceRepo interface {
	Create(device structures.Device) (int, error)
	GetAll() ([]structures.Device, error)
	Get(id int) (structures.Device, error)
	Update(id int, input structures.UpdateDeviceInput) error
	Delete(id int) error
}

type NotificationRepo interface {
	Create(notification structures.Notification) (int, error)
	GetAll() ([]structures.Notification, error)
	Get(id int) (structures.Notification, error)
	GetAllByPatientID(patientID int) ([]structures.Notification, error)
}

type IndicatorsStampRepo interface {
	Create(indicatorsStamp structures.IndicatorsStamp) (int, error)
	GetAll() ([]structures.IndicatorsStamp, error)
	GetById(id int) (structures.IndicatorsStamp, error)
}

type Repository struct {
	AuthorizationRepo
	UserRepo
	PatientRepo
	MedicineRepo
	DiagnosisRepo
	PatientMedicineRepo
	UserPatientRepo
	DeviceRepo
	NotificationRepo
	IndicatorsStampRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepo:   NewAuthPostgres(db),
		UserRepo:            NewUserActionPostgres(db),
		PatientRepo:         NewPatientPostgres(db),
		MedicineRepo:        NewMedicinePostgres(db),
		DiagnosisRepo:       NewDiagnosisPostgres(db),
		PatientMedicineRepo: NewPatientMedicinePostgres(db),
		UserPatientRepo:     NewUserPatientPostgres(db),
		DeviceRepo:          NewDevicePostgres(db),
		NotificationRepo:    NewNotificationPostgres(db),
		IndicatorsStampRepo: NewIndicatorsStampPostgres(db),
	}
}
