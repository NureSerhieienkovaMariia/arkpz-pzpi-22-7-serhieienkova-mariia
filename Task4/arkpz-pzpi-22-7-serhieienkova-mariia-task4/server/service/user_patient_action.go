package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type UserPatientActionService struct {
	repo repository.UserPatientRepo
}

func NewUserPatientActionService(repo repository.UserPatientRepo) *UserPatientActionService {
	return &UserPatientActionService{repo: repo}
}

func (s *UserPatientActionService) Create(userPatient structures.UserPatient) (int, error) {
	return s.repo.Create(userPatient)
}

func (s *UserPatientActionService) GetAll() ([]structures.UserPatient, error) {
	return s.repo.GetAll()
}

func (s *UserPatientActionService) Get(id int) (structures.UserPatient, error) {
	return s.repo.Get(id)
}

func (s *UserPatientActionService) Update(id int, input structures.UserPatient) error {
	return s.repo.Update(id, input)
}

func (s *UserPatientActionService) Delete(id int) error {
	return s.repo.Delete(id)
}
