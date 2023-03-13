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

// GetAll implements participant.ServiceInterface
func (s *participantService) GetAll(page int, limit int) (data []participant.ParticipantCore, totalPage int, err error) {
	offset := (page - 1) * limit
	dataParticipant, count, errParticipant := s.participantRepository.GetAll(limit, offset)
	if errParticipant != nil {
		err = errParticipant
		return nil, 0, err
	}

	if count < 10 {
		totalPage = 1
	} else if int(count)%limit == 0 {
		totalPage = int(count) / limit
	} else {
		totalPage = (int(count) / limit) + 1
	}

	data = dataParticipant

	return

}
