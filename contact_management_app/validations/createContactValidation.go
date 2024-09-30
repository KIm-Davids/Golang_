package validations

import (
	"fmt"
	"log"
	request2 "project/dtos/request"
	"project/repository"
	"regexp"
)

func ValidatePhoneNumberLength(request request2.CreateContactRequest) bool {

	var regexPatter, err = regexp.Compile(`^\d{11}$`)
	if err != nil {
		log.Fatalf("An error occurred in validating create contact phone number")
	}
	if regexPatter.MatchString(request.PhoneNumber) {
		return true
	}
	return false
}

func ValidateContactEmail(request request2.CreateContactRequest) bool {
	var regex, err = regexp.Compile(`^[a-zA-Z0-9+/.-_%&$]+[@a-zA-Z.-]+\.[a-z]{2,}$`)

	if err != nil {
		log.Fatal("An error occurred in verifying email")
	}

	if regex.MatchString(request.Email) {
		return true
	}
	return false
}

func CheckForDuplicatesInTheBackEnd(request request2.CreateContactRequest) bool {
	contact := repository.RetrieveContactFromContactDbForValidation()

	fmt.Print(contact)
	fmt.Print(request)
	//if request.Email == contact && request.PhoneNumber == contact.PhoneNumber {
	//	return false
	//} else {
	return false

}
