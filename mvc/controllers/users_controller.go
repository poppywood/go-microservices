package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/poppywood/go-microservices/mvc/services"
	"github.com/poppywood/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appError := &utils.ApplicationError{
			Message:    fmt.Sprintf("user_id %v is not a number", userId),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(appError)
		resp.WriteHeader(appError.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, appError := services.GetUser(userId)
	if appError != nil {
		jsonValue, _ := json.Marshal(appError)
		resp.WriteHeader(appError.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
