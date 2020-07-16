package services

import (
	"github.com/poppywood/go-microservices/mvc/domain"
	"github.com/poppywood/go-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
