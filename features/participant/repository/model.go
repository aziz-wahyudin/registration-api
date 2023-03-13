package repository

import (
	"github.com/aziz-wahyudin/registration-api/features/participant"
	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	Name  string
	Phone string `gorm:"unique"`
	Age   int
	Photo string
}

func FromCore(dataCore participant.ParticipantCore) Participant {
	userGorm := Participant{
		Name:  dataCore.Name,
		Phone: dataCore.Phone,
		Age:   dataCore.Age,
		Photo: dataCore.Photo,
	}
	return userGorm
}

func (dataModel *Participant) ToCore() participant.ParticipantCore {
	return participant.ParticipantCore{
		Id:    dataModel.ID,
		Name:  dataModel.Name,
		Phone: dataModel.Phone,
		Photo: dataModel.Photo,
		Age:   dataModel.Age,
	}
}

func ListToCore(dataModel []Participant) []participant.ParticipantCore {
	var dataCore []participant.ParticipantCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}
