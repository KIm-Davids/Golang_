package web

import (
	"encoding/json"
	"net/http"
	request2 "project/dtos/request"
	response2 "project/dtos/response"
	"project/service"
)

func RegisterUser(registerUserResponse http.ResponseWriter, registerUserRequest *http.Request) {
	response, err := service.RegisterUserServices(registerUserRequest)
	if err != nil {
		registerUserResponse.WriteHeader(http.StatusBadRequest)
	}
	if data, ok := response.(map[string]string); ok {
		json.NewEncoder(registerUserResponse).Encode(response2.ApiResponse{Success: true, Data: data})
	} else {
		json.NewEncoder(registerUserResponse).Encode(response2.ApiResponse{Data: response})
	}
}

func LoginUser(loginUserResponse http.ResponseWriter, loginUserRequest *http.Request) {
	var request request2.LoginRequest

	response, err := service.LoginUserServices(loginUserRequest)

}
