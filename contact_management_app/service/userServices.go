package service

import (
	"encoding/json"
	"net/http"
	request2 "project/dtos/request"
	"project/validations"
)

func RegisterUserServices(registerUserRequest *http.Request) (interface{}, error) {
	var request request2.RegisterUserRequest
	response, err := validations.RegisterUserValidation(request)
	if err != nil {
		return nil, nil
	}
	if response != nil {
		return response, nil
	}

	err = json.NewDecoder(registerUserRequest.Body).Decode(&request)
	if err != nil {
		response := map[string]string{
			"message": "An error occurred",
		}
		return response, error(http.ErrUseLastResponse)

	} else {
		response := map[string]string{
			"message": "User Registered Successfully",
		}
		return response, nil
	}
}

func LoginUserServices(loginUserRequest *http.Request) (interface{}, error) {
	var request request2.LoginRequest
	err := json.NewDecoder(loginUserRequest.Body).Decode(&request)

	if err != nil {
		return err, nil
	}
	return nil, nil
}
