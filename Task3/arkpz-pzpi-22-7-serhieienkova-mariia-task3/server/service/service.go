package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type Authorization interface {
	CreateUser(user structures.User) (int, error)
	GenerateToken(username, password string) (structures.UserToken, error)
	GenerateTokenByUserId(userId int) (structures.UserToken, error)
	RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error)
	GetUserById(userId int) (structures.User, error)
}

type UserAction interface {
	CreateUser(user structures.User) (int, error)
	GetAll() ([]structures.User, error)
	GetById(userId int) (structures.User, error)
	Delete(userId int) error
	Update(userId int, updatedUser structures.UpdateUserInput) error
	GetByEmail(email string) (structures.User, error)
}

type PatientAction interface {
	Create(patient structures.CreatePatientInput) (int, error)
	GetAll() ([]structures.Patient, error)
	GetById(id int) (structures.Patient, error)
	Update(id int, input structures.UpdatePatientInput) error
	Delete(id int) error
	GetFullInfo(patientID int) (structures.PatientFullInfo, error)
}

type MedicineAction interface {
	Create(medicine structures.Medicine) (int, error)
	GetAll() ([]structures.Medicine, error)
	Get(id int) (structures.Medicine, error)
	Update(id int, input structures.UpdateMedicineInput) error
	Delete(id int) error
}

type DiagnosisAction interface {
	Create(diagnosis structures.Diagnosis) (int, error)
	GetAll() ([]structures.Diagnosis, error)
	Get(id int) (structures.Diagnosis, error)
	Update(id int, input structures.UpdateDiagnosisInput) error
	Delete(id int) error
}

type PatientMedicineAction interface {
	Create(patientMedicine structures.PatientMedicine) (int, error)
	GetAll() ([]structures.PatientMedicine, error)
	Get(id int) (structures.PatientMedicine, error)
	Update(id int, input structures.PatientMedicine) error
	Delete(id int) error
	SetMedicineToPatient(patientID, medicineID int, schedule string) error
}

type UserPatientAction interface {
	Create(userPatient structures.UserPatient) (int, error)
	GetAll() ([]structures.UserPatient, error)
	Get(id int) (structures.UserPatient, error)
	Update(id int, input structures.UserPatient) error
	Delete(id int) error
}

type DeviceAction interface {
	Create(device structures.Device) (int, error)
	GetAll() ([]structures.Device, error)
	Get(id int) (structures.Device, error)
	Update(id int, input structures.UpdateDeviceInput) error
	Delete(id int) error
}

type NotificationAction interface {
	Create(notification structures.Notification) (int, error)
	GetAll() ([]structures.Notification, error)
	Get(id int) (structures.Notification, error)
	GetAllByPatientID(patientID int) ([]structures.Notification, error)
}

type IndicatorsStampAction interface {
	Create(input structures.IndicatorsStamp) error
	GetAll() ([]structures.IndicatorsStamp, error)
	GetById(id int) (structures.IndicatorsStamp, error)
}

type Service struct {
	Authorization
	UserAction
	PatientAction
	MedicineAction
	DiagnosisAction
	PatientMedicineAction
	UserPatientAction
	DeviceAction
	NotificationAction
	IndicatorsStampAction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:         NewAuthService(repos.AuthorizationRepo),
		UserAction:            NewUserActionService(repos.UserRepo),
		PatientAction:         NewPatientActionService(repos.PatientRepo),
		MedicineAction:        NewMedicineActionService(repos.MedicineRepo),
		DiagnosisAction:       NewDiagnosisActionService(repos.DiagnosisRepo),
		PatientMedicineAction: NewPatientMedicineActionService(repos.PatientMedicineRepo),
		UserPatientAction:     NewUserPatientActionService(repos.UserPatientRepo),
		DeviceAction:          NewDeviceActionService(repos.DeviceRepo),
		NotificationAction:    NewNotificationActionService(repos.NotificationRepo),
		IndicatorsStampAction: NewIndicatorsStampActionService(repos.IndicatorsStampRepo, repos.NotificationRepo),
	}
}
