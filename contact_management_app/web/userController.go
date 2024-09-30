package web

import (
	"encoding/json"
	"log"
	"net/http"
	response2 "project/dtos/response"
	"project/service"
)

func RegisterUser(registerUserResponse http.ResponseWriter, registerUserRequest *http.Request) {

	if !(registerUserRequest.Method == http.MethodPost) {
		registerUserResponse.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response, err := service.RegisterUserServices(registerUserRequest)

	if err != nil {
		registerUserResponse.WriteHeader(http.StatusBadRequest)
	}

	if _, ok := response.(map[string]string); ok {
		registerUserResponse.WriteHeader(http.StatusAccepted)
	}
	if response == true {
		json.NewEncoder(registerUserResponse).Encode(response2.ApiResponse{Success: true, Data: "User Registered Successfully"})
	}

	if response == false {
		registerUserResponse.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(registerUserResponse).Encode(response2.ApiResponse{Success: false, Data: "Please ensure your details are valid"})
	}
}

func LoginUser(loginUserResponse http.ResponseWriter, loginUserRequest *http.Request) {

	if !(loginUserRequest.Method == http.MethodPatch) {
		loginUserResponse.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response, err := service.LoginUserServices(loginUserRequest)
	if err != nil {
		log.Fatalf("An error occurred %v", err)
	}

	if _, ok := response.(map[string]string); ok {
		loginUserResponse.WriteHeader(http.StatusAccepted)
		json.NewEncoder(loginUserResponse).Encode(response2.ApiResponse{Success: true, Data: "User Logged in successfully"})
	}

	if response == false {
		loginUserResponse.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(loginUserResponse).Encode(response2.ApiResponse{Success: false, Data: "An error occurred, Please try again "})
	}

}

func CreateContact(createContactResponse http.ResponseWriter, createContactRequest *http.Request) {

	if !(createContactRequest.Method == http.MethodPatch) {
		createContactResponse.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response := service.CreateContactService(createContactRequest)

	if response == true {
		json.NewEncoder(createContactResponse).Encode(response2.ApiResponse{Success: true, Data: "Added a contact Successfully"})
	}

	if response == false {
		createContactResponse.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(createContactResponse).Encode(response2.ApiResponse{Success: false, Data: "Something went wrong, Please check to see if your entry is valid "})
	}

}

func GetAllContacts(response http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	allContacts, err := service.ViewAllContacts()

	if allContacts != nil {
		response.WriteHeader(http.StatusAccepted)
		json.NewEncoder(response).Encode(response2.ApiResponse{Success: true, Data: allContacts})
	}

	if err != nil {
		response.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(response).Encode(response2.ApiResponse{Success: false, Data: "Error retrieving all contacts"})
	}
}
