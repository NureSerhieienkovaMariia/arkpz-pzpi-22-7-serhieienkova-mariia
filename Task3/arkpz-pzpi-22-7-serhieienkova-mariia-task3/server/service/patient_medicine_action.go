package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"time"
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

func (s *PatientMedicineActionService) SetMedicineToPatient(patientID, medicineID int, schedule string) error {
	patientMedicine := structures.PatientMedicine{
		PatientId:  patientID,
		MedicineId: medicineID,
		Date:       time.Now().Format(time.RFC3339),
		Schedule:   schedule,
	}
	return s.repo.SetMedicineToPatient(patientMedicine)
}
