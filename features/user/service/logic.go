package service

import (
	"errors"

	"github.com/aziz-wahyudin/registration-api/features/user"
	"github.com/aziz-wahyudin/registration-api/middlewares"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (s *userService) Create(input user.UserCore) (err error) {
	input.Role = "participant"
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := s.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// Login implements user.ServiceInterface
func (s *userService) Login(email string, password string) (data user.UserCore, token string, err error) {
	data, err = s.userRepository.FindUser(email)
	if err != nil {
		return user.UserCore{}, "", errors.New("failed login")
	}

	token, err = middlewares.CreateToken(int(data.Id), data.Role)
	if err != nil {
		return user.UserCore{}, "", errors.New("failed login")
	}

	return data, token, nil
}
