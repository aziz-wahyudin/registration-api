package delivery

import (
	"github.com/aziz-wahyudin/registration-api/features/user"
)

type UserReq struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func toCore(data UserReq) user.UserCore {
	return user.UserCore{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
