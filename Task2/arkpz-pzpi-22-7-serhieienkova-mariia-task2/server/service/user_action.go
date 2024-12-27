package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type UserActionService struct {
	repo repository.UserRepo
}

func (s *UserActionService) CreateUser(user structures.User) (int, error) {
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	return s.repo.CreateUser(user)
}

func NewUserActionService(repo repository.UserRepo) *UserActionService {
	return &UserActionService{repo: repo}
}

func (s *UserActionService) GetAll() ([]structures.User, error) {
	users, err := s.repo.GetAll()
	return users, err
}

func (s *UserActionService) GetById(userId int) (structures.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserActionService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserActionService) Update(userId int, input structures.UpdateUserInput) error {
	if input.PasswordHash != "" {
		input.PasswordHash = generatePasswordHash(input.PasswordHash)
	}

	return s.repo.Update(userId, input)
}

func (s *UserActionService) GetByEmail(email string) (structures.User, error) {
	return s.repo.GetByEmail(email)
}
