package user

type UserService interface {
	GetWelcomeMessage() string
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetWelcomeMessage() string {
	return "Hello from UserService!"
}
