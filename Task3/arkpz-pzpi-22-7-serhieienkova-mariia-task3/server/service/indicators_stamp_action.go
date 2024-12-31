package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
	"fmt"
	"time"
)

type IndicatorsStampActionService struct {
	repo             repository.IndicatorsStampRepo
	notificationRepo repository.NotificationRepo
}

func NewIndicatorsStampActionService(repo repository.IndicatorsStampRepo, notificationRepo repository.NotificationRepo) *IndicatorsStampActionService {
	return &IndicatorsStampActionService{repo: repo, notificationRepo: notificationRepo}
}

func (s *IndicatorsStampActionService) Create(input structures.IndicatorsStamp) error {
	input.Timestamp = time.Now().Format(time.RFC3339)
	id, err := s.repo.Create(input)
	if err != nil {
		return err
	}

	fmt.Println("IndicatorsStampActionService.Create: ", input)
	message := s.checkIndicators(input)
	fmt.Println("IndicatorsStampActionService.Create message: ", message)
	if message != "" {
		notification := structures.Notification{
			IndicatorStampId: id,
			Message:          message,
			Timestamp:        input.Timestamp,
		}
		_, err := s.notificationRepo.Create(notification)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *IndicatorsStampActionService) checkIndicators(stamp structures.IndicatorsStamp) string {
	var message string

	if stamp.Pulse < 60 || stamp.Pulse > 100 {
		if stamp.Pulse < 50 || stamp.Pulse > 110 {
			message += "Critical heart rate detected. "
		} else {
			message += "Warning: Abnormal heart rate detected. "
		}
	}
	if stamp.SystolicBloodPressure < 90 || stamp.SystolicBloodPressure > 120 {
		if stamp.SystolicBloodPressure < 80 || stamp.SystolicBloodPressure > 130 {
			message += "Critical systolic blood pressure detected. "
		} else {
			message += "Warning: Abnormal systolic blood pressure detected. "
		}
	}
	if stamp.DiastolicBloodPressure < 60 || stamp.DiastolicBloodPressure > 80 {
		if stamp.DiastolicBloodPressure < 50 || stamp.DiastolicBloodPressure > 90 {
			message += "Critical diastolic blood pressure detected. "
		} else {
			message += "Warning: Abnormal diastolic blood pressure detected. "
		}
	}
	if stamp.Temperature < 36.1 || stamp.Temperature > 37.2 {
		if stamp.Temperature < 35.5 || stamp.Temperature > 38.0 {
			message += "Critical temperature detected. "
		} else {
			message += "Warning: Abnormal temperature detected. "
		}
	}

	return message
}
func (s *IndicatorsStampActionService) GetAll() ([]structures.IndicatorsStamp, error) {
	return s.repo.GetAll()
}

func (s *IndicatorsStampActionService) GetById(id int) (structures.IndicatorsStamp, error) {
	return s.repo.GetById(id)
}
