package delivery

import "github.com/aziz-wahyudin/registration-api/features/participant"

type ParticipatResp struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
	Photo string `json:"photo"`
}

func fromCore(dataCore participant.ParticipantCore) ParticipatResp {
	return ParticipatResp{
		Name:  dataCore.Name,
		Phone: dataCore.Phone,
		Age:   dataCore.Age,
		Photo: dataCore.Photo,
	}
}

func FromCoreList(dataCore []participant.ParticipantCore) []ParticipatResp {
	var dataResp []ParticipatResp
	for _, v := range dataCore {
		dataResp = append(dataResp, fromCore(v))
	}
	return dataResp
}
