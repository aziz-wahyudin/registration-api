package repository

import (
	"errors"

	"github.com/aziz-wahyudin/registration-api/features/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.RepositoryInterface
func (r *userRepository) Create(input user.UserCore) (row int, err error) {
	userGorm := FromCore(input)
	result := r.db.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)", userGorm.Name, userGorm.Email, userGorm.Password, userGorm.Role)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(result.RowsAffected), nil
}

// FindUser implements user.RepositoryInterface
func (r *userRepository) FindUser(email string) (result user.UserCore, err error) {
	var userData User
	tx := r.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return user.UserCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return user.UserCore{}, errors.New("login failed")
	}

	result = userData.ToCore()

	return result, nil
}
