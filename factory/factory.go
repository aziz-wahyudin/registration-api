package factory

import (
	userDelivery "github.com/aziz-wahyudin/registration-api/features/user/delivery"
	userRepo "github.com/aziz-wahyudin/registration-api/features/user/repository"
	userService "github.com/aziz-wahyudin/registration-api/features/user/service"

	participantDelivery "github.com/aziz-wahyudin/registration-api/features/participant/delivery"
	participantRepo "github.com/aziz-wahyudin/registration-api/features/participant/repository"
	participantService "github.com/aziz-wahyudin/registration-api/features/participant/service"

	// "net/http"

	"gorm.io/gorm"
)

func InitFactory(db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory)

	participantRepoFactory := participantRepo.New(db)
	participantServiceFactory := participantService.New(participantRepoFactory)
	participantDelivery.New(participantServiceFactory)
}
