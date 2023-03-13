package participant

type ParticipantCore struct {
	Id    uint
	Name  string
	Phone string `gorm:"unique"`
	Age   int
	Photo string
}

type ServiceInterface interface {
	Create(input ParticipantCore) (err error)
	Update(input ParticipantCore, id uint) error
	GetAll(page, limit int) (data []ParticipantCore, totalPage int, err error)
}

type RepositoryInterface interface {
	Create(input ParticipantCore) (row int, err error)
	Update(input ParticipantCore, id uint) error
	GetAll(limit, offset int) (data []ParticipantCore, count int64, err error)
}
