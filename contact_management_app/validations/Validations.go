package validations

import (
	request2 "project/dtos/request"
	"project/exceptions"
)

//
//func validateContacts(details) error {
//
//	if !strings.Contains(deta, "@mail.com") {
//		return &exceptions.InvalidDetailsException{Message: "Invalid Details, Please try again"}
//	}
//
//}

func RegisterUserValidation(request request2.RegisterUserRequest) (interface{}, error) {

	if request.FirstName == " " || request.LastName == " " || request.Email == " " || request.PhoneNumber == " " {
		return exceptions.InvalidDetailsException{Message: "Please fill in all the details !"}, nil
	}
	return nil, nil

}
