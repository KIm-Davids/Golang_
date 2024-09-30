package validations

import (
	"log"
	request2 "project/dtos/request"
	"project/repository"
	"regexp"
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

func CheckForUserInTheDb(request request2.RegisterUserRequest) interface{} {
	users := repository.RetrieveValuesFromTheDb()

	for _, values := range users {
		if request.Email == values.Email || request.PhoneNumber == values.PhoneNumber {
			return false
		}
	}
	return true
}

func CheckPhoneNumberLength(request request2.RegisterUserRequest) bool {

	var regexPattern, err = regexp.Compile(`^\d{11}$`)
	if err != nil {
		log.Fatalf("An error occurred with the regex %v", err)
	}

	if regexPattern.MatchString(request.PhoneNumber) {
		return true
	}
	return false

}
