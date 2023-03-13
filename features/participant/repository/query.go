package repository

import (
	"errors"

	"github.com/aziz-wahyudin/registration-api/features/participant"

	"gorm.io/gorm"
)

type participantRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) participant.RepositoryInterface {
	return &participantRepository{
		db: db,
	}
}

// Create implements participant.RepositoryInterface
func (r *participantRepository) Create(input participant.ParticipantCore) (row int, err error) {
	userGorm := FromCore(input)
	result := r.db.Exec("INSERT INTO participants (name, phone, age, photo) VALUES (?, ?, ?, ?)", userGorm.Name, userGorm.Phone, userGorm.Age, userGorm.Photo)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(result.RowsAffected), nil
}
