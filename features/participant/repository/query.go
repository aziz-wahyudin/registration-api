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
	participantGorm := FromCore(input)
	result := r.db.Exec("INSERT INTO participants (name, phone, age, photo) VALUES (?, ?, ?, ?)", participantGorm.Name, participantGorm.Phone, participantGorm.Age, participantGorm.Photo)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(result.RowsAffected), nil
}

// Update implements participant.RepositoryInterface
func (r *participantRepository) Update(input participant.ParticipantCore, id uint) error {
	participantGorm := FromCore(input)
	result := r.db.Exec("UPDATE participants SET name = ?, phone = ?, age = ?, photo = ?, updated_at = NOW() WHERE id = ?", participantGorm.Name, participantGorm.Phone, participantGorm.Age, participantGorm.Photo, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetAll implements participant.RepositoryInterface
func (r *participantRepository) GetAll(limit int, offset int) (data []participant.ParticipantCore, count int64, err error) {
	var participants []Participant

	query := "SELECT * FROM participants"
	tx := r.db.Raw(query).Scan(&participants).Count(&count)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}

	tx1 := r.db.Raw(query).Limit(limit).Offset(offset).Find(&participants)
	if tx1.Error != nil {
		return nil, 0, tx1.Error
	}
	if tx1.RowsAffected == 0 {
		return nil, 0, errors.New("get all data failed, error query data")
	}

	data = ListToCore(participants)
	return data, count, nil
}
