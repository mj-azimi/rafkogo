package landing

type LandingService interface {
	GetWelcomeMessage() string
}

type landingService struct{}

func NewLandingService() LandingService {
	return &landingService{}
}

func (s *landingService) GetWelcomeMessage() string {
	return "Hello from UserService!"
}
