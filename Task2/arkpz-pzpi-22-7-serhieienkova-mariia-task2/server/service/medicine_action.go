package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type MedicineActionService struct {
	repo repository.MedicineRepo
}

func NewMedicineActionService(repo repository.MedicineRepo) *MedicineActionService {
	return &MedicineActionService{repo: repo}
}

func (s *MedicineActionService) Create(medicine structures.Medicine) (int, error) {
	return s.repo.Create(medicine)
}

func (s *MedicineActionService) GetAll() ([]structures.Medicine, error) {
	return s.repo.GetAll()
}

func (s *MedicineActionService) Get(id int) (structures.Medicine, error) {
	return s.repo.Get(id)
}

func (s *MedicineActionService) Update(id int, input structures.Medicine) error {
	return s.repo.Update(id, input)
}

func (s *MedicineActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
