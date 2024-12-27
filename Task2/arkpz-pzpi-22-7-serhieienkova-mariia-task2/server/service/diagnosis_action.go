package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type DiagnosisActionService struct {
	repo repository.DiagnosisRepo
}

func NewDiagnosisActionService(repo repository.DiagnosisRepo) *DiagnosisActionService {
	return &DiagnosisActionService{repo: repo}
}

func (s *DiagnosisActionService) Create(diagnosis structures.Diagnosis) (int, error) {
	return s.repo.Create(diagnosis)
}

func (s *DiagnosisActionService) GetAll() ([]structures.Diagnosis, error) {
	return s.repo.GetAll()
}

func (s *DiagnosisActionService) Get(id int) (structures.Diagnosis, error) {
	return s.repo.Get(id)
}

func (s *DiagnosisActionService) Update(id int, input structures.Diagnosis) error {
	return s.repo.Update(id, input)
}

func (s *DiagnosisActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
