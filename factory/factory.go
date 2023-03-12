package factory

import (
	userDelivery "github.com/aziz-wahyudin/registration-api/features/user/delivery"
	userRepo "github.com/aziz-wahyudin/registration-api/features/user/repository"
	userService "github.com/aziz-wahyudin/registration-api/features/user/service"

	// "net/http"

	"gorm.io/gorm"
)

func InitFactory(db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory)
}
