package service

import (
	"errors"

	"github.com/aziz-wahyudin/registration-api/features/participant"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type participantService struct {
	participantRepository participant.RepositoryInterface
	validate              *validator.Validate
}

func New(repo participant.RepositoryInterface) participant.ServiceInterface {
	return &participantService{
		participantRepository: repo,
		validate:              validator.New(),
	}
}

// Create implements participant.ServiceInterface
func (s *participantService) Create(input participant.ParticipantCore) (err error) {
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := s.participantRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// Update implements participant.ServiceInterface
func (s *participantService) Update(input participant.ParticipantCore, id uint) error {
	err := s.participantRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
