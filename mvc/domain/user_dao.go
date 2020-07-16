package domain

import (
	"fmt"
	"github.com/poppywood/go-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		123: &User{Id: 123, FirstName: "Ann", LastName: "Doe", Email: "some@thing.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
