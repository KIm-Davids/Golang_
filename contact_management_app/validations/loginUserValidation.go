package validations

import (
	"project/dtos/request"
	"strings"
)

func CheckForEmptySpace(loginRequest request.LoginRequest) bool {

	if strings.Contains(loginRequest.Email, " ") {
		return false
	}
	return true
}

func CheckForCorrectEmail(loginRequest request.LoginRequest) bool {

	if !(strings.Contains(loginRequest.Email, "@gmail.com")) {
		return false
	}
	return true
}

//func CheckForCorrectPassword(loginRequest request.LoginRequest) bool {
//
//	password := repository.RetrievePassword()
//
//	if loginRequest.Password == password {
//		return true
//	}
//	return false
//}
