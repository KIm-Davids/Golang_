package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	request2 "project/dtos/request"
	"project/models"
	"project/repository"
	"project/validations"
)

func CreateContactService(createContactRequest *http.Request) interface{} {
	var request request2.CreateContactRequest

	err := json.NewDecoder(createContactRequest.Body).Decode(&request)

	userDoesNotExistsInTheDb := validations.CheckForDuplicatesInTheBackEnd(request)
	validatePhoneNumber := validations.ValidatePhoneNumberLength(request)
	emailResponse := validations.ValidateContactEmail(request)

	//loginStatus := repository.GetLoginStatusOfUser()

	fmt.Print(userDoesNotExistsInTheDb)
	if emailResponse && validatePhoneNumber && userDoesNotExistsInTheDb {
		response := repository.SaveContactToTheDb(request)
		if response == true {
			return true
		}
	}

	if emailResponse == false || validatePhoneNumber == false || userDoesNotExistsInTheDb == false {
		return false
	}

	if err != nil {
		log.Fatalf("An error occurred in decoding the request %v", err)
	}

	return nil
}

func ViewAllContacts() ([]models.ContactDetails, error) {
	allContacts, err := repository.RetrieveAllValuesForUser()

	fmt.Print(allContacts)

	return allContacts, err
}
