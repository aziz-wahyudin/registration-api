package delivery

import (
	"github.com/aziz-wahyudin/registration-api/features/participant"
)

type ParticipantReq struct {
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Age   int    `json:"age" form:"age"`
	Photo string `json:"photo" form:"photo"`
}

func toCore(data ParticipantReq) participant.ParticipantCore {
	return participant.ParticipantCore{
		Name:  data.Name,
		Phone: data.Phone,
		Age:   data.Age,
		Photo: data.Photo,
	}
}
