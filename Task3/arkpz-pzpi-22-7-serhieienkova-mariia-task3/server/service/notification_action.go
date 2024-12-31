package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type NotificationActionService struct {
	repo repository.NotificationRepo
}

func NewNotificationActionService(repo repository.NotificationRepo) *NotificationActionService {
	return &NotificationActionService{repo: repo}
}

func (s *NotificationActionService) Create(notification structures.Notification) (int, error) {
	return s.repo.Create(notification)
}

func (s *NotificationActionService) GetAll() ([]structures.Notification, error) {
	return s.repo.GetAll()
}

func (s *NotificationActionService) Get(id int) (structures.Notification, error) {
	return s.repo.Get(id)
}

func (s *NotificationActionService) GetAllByPatientID(patientID int) ([]structures.Notification, error) {
	return s.repo.GetAllByPatientID(patientID)
}
