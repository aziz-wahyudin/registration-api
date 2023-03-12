package repository

import (
	"github.com/aziz-wahyudin/registration-api/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

func FromCore(dataCore user.UserCore) User {
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Role:     dataCore.Role,
	}
	return userGorm
}

func (dataModel *User) ToCore() user.UserCore {
	return user.UserCore{
		Id:    dataModel.ID,
		Name:  dataModel.Name,
		Email: dataModel.Email,
		Role:  dataModel.Role,
	}
}

func ListToCore(dataModel []User) []user.UserCore {
	var dataCore []user.UserCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}
