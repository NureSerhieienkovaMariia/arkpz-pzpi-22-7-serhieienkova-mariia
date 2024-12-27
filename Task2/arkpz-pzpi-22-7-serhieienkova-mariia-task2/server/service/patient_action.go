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

func (s *PatientActionService) Create(patient structures.Patient) (int, error) {
	return s.repo.Create(patient)
}

func (s *PatientActionService) GetAll() ([]structures.Patient, error) {
	return s.repo.GetAll()
}

func (s *PatientActionService) GetById(id int) (structures.Patient, error) {
	return s.repo.GetById(id)
}

func (s *PatientActionService) Update(id int, input structures.Patient) error {
	return s.repo.Update(id, input)
}

func (s *PatientActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
