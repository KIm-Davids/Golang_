package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	request2 "project/dtos/request"
	"project/repository"
	"project/validations"
)

func RegisterUserServices(registerUserRequest *http.Request) (interface{}, interface{}) {
	var request request2.RegisterUserRequest

	err := json.NewDecoder(registerUserRequest.Body).Decode(&request)

	noWhiteSpaceResponse := validations.CheckForWhiteSpaces(request)
	notAUserInTheDB := validations.CheckForUserInTheDb(request)
	emailValidation := validations.EmailValidation(request)
	checkPhoneNumberLength := validations.CheckPhoneNumberLength(request)

	fmt.Print(notAUserInTheDB)

	if checkPhoneNumberLength == true && noWhiteSpaceResponse == true && notAUserInTheDB == true && emailValidation == true {
		repository.InsertUsersIntoDatabase(request.FirstName, request.LastName, request.Email, request.PhoneNumber, request.Password)
		return true, nil
	}

	if checkPhoneNumberLength == false || noWhiteSpaceResponse == false || notAUserInTheDB == false || emailValidation == false {
		return false, nil
	}

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

	noEmptySpace := validations.CheckForEmptySpace(request)
	correctEmail := validations.CheckForCorrectEmail(request)
	//correctPassword := validations.CheckForCorrectPassword(request)
	users := repository.RetrieveValuesFromTheDb()

	if noEmptySpace && correctEmail {
		for _, user := range users {
			if request.Email == user.Email && request.Password == user.Password {
				user.LoginStatus = "online"
				repository.SetLoginStatus(user.LoginStatus, request.Email, request.Password)

				response := map[string]string{
					"message": "User Logged in successfully",
				}
				return response, nil
			}
			if request.Password != user.Password {
				return false, nil
			}
		}
		return false, nil
	}

	if noEmptySpace == false && correctEmail == false {
		return false, nil
	}

	if err != nil {
		log.Fatalf("An error occurred can't login")
	}
	return nil, nil
}
