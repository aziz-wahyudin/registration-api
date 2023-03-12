package user

type UserCore struct {
	Id       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string
}

type ServiceInterface interface {
	Create(input UserCore) (err error)
	Login(email string, password string) (data UserCore, token string, err error)
}

type RepositoryInterface interface {
	Create(input UserCore) (row int, err error)
	FindUser(email string) (result UserCore, err error)
}
