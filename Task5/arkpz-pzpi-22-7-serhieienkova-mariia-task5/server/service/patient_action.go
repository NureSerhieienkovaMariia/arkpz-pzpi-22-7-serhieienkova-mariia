package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type PatientActionService struct {
	repo repository.PatientRepo
}

func NewPatientActionService(repo repository.PatientRepo) *PatientActionService {
	return &PatientActionService{repo: repo}
}

func (s *PatientActionService) Create(input structures.CreatePatientInput) (int, error) {
	patient := structures.Patient{
		Name:        input.Name,
		Surname:     input.Surname,
		Birthday:    input.Birthday,
		DiagnosisId: input.DiagnosisId,
	}
	return s.repo.Create(patient)
}

func (s *PatientActionService) Update(id int, input structures.UpdatePatientInput) error {
	return s.repo.Update(id, input)
}

func (s *PatientActionService) GetAll() ([]structures.Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientActionService) GetById(id int) (structures.Patient, error) {
	return s.repo.GetById(id)
}

func (s *PatientActionService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PatientActionService) GetFullInfo(patientID int) (structures.PatientFullInfo, error) {
	patient, err := s.repo.GetById(patientID)
	if err != nil {
		return structures.PatientFullInfo{}, err
	}

	diagnoses, err := s.repo.GetDiagnosesByPatientID(patientID)
	if err != nil {
		return structures.PatientFullInfo{}, err
	}

	medicines, err := s.repo.GetMedicinesByPatientID(patientID)
	if err != nil {
		return structures.PatientFullInfo{}, err
	}

	devices, err := s.repo.GetDevicesByPatientID(patientID)
	if err != nil {
		return structures.PatientFullInfo{}, err
	}

	indicators, err := s.repo.GetIndicatorsByPatientID(patientID)
	if err != nil {
		return structures.PatientFullInfo{}, err
	}

	notifications, err := s.repo.GetNotificationsByPatientID(patientID)
	if err != nil {
		return structures.PatientFullInfo{}, err
	}

	return structures.PatientFullInfo{
		Patient:       patient,
		Diagnoses:     diagnoses,
		Medicines:     medicines,
		Devices:       devices,
		Indicators:    indicators,
		Notifications: notifications,
	}, nil
}
