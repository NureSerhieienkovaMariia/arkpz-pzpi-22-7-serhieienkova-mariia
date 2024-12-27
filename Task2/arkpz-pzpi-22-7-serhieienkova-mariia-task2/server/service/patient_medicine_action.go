package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type PatientMedicineActionService struct {
	repo repository.PatientMedicineRepo
}

func NewPatientMedicineActionService(repo repository.PatientMedicineRepo) *PatientMedicineActionService {
	return &PatientMedicineActionService{repo: repo}
}

func (s *PatientMedicineActionService) Create(patientMedicine structures.PatientMedicine) (int, error) {
	return s.repo.Create(patientMedicine)
}

func (s *PatientMedicineActionService) GetAll() ([]structures.PatientMedicine, error) {
	return s.repo.GetAll()
}

func (s *PatientMedicineActionService) Get(id int) (structures.PatientMedicine, error) {
	return s.repo.Get(id)
}

func (s *PatientMedicineActionService) Update(id int, input structures.PatientMedicine) error {
	return s.repo.Update(id, input)
}

func (s *PatientMedicineActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
