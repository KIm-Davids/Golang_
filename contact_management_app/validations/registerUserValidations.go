package validations

import (
	request2 "project/dtos/request"
	"project/repository"
	"strings"
)

func CheckForWhiteSpaces(request request2.RegisterUserRequest) bool {

	if strings.ContainsAny(request.FirstName, " ") || strings.ContainsAny(request.LastName, " ") || strings.ContainsAny(request.Email, " ") || strings.ContainsAny(request.PhoneNumber, " ") {
		return false
	} else {
		return true
	}

}

func EmailValidation(request request2.RegisterUserRequest) bool {
	if strings.Contains(request.Email, "@email.com") || strings.Contains(request.Email, "@gmail.com") {
		return true
	} else {
		return false
	}
}

func CheckForUserInTheDb(request request2.RegisterUserRequest) bool {
	users := repository.RetrieveValuesFromTheDb()

	for _, values := range users {
		if request.Email == values.Email || request.PhoneNumber == values.PhoneNumber {
			return false
		}
	}
	return true
}

func CheckPhoneNumberLength(request request2.RegisterUserRequest) bool {

	if len(request.PhoneNumber) != 11 {
		return false
	}
	return true
}
