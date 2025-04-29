package landing

import (
	landingModel "rafkogo/internal/landing/database/landing"
	"rafkogo/utils/mysql"
)
type LandingService interface {
	GetWelcomeMessage() string
	GetFirstLanding() (*string, error)
	GetAllLandings() ([]landingModel.Landing, error)
}

type landingService struct{}

func NewLandingService() LandingService {
	return &landingService{}
}

func (s *landingService) GetWelcomeMessage() string {
	return "Hello from UserService!"
}


func (s *landingService) GetFirstLanding() (*string, error) {
	var landing landingModel.Landing
	result := mysql.DB.First(&landing)
	if result.Error != nil {
		return nil, result.Error
	}
	return landing.Title, nil
}

func (s *landingService) GetAllLandings() ([]landingModel.Landing, error) {
	var landings []landingModel.Landing
	result := mysql.DB.Find(&landings)
	if result.Error != nil {
		return nil, result.Error
	}
	return landings, nil
}