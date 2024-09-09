package request

type RegisterUserRequest struct {
	FirstName   string `json: "firstName"`
	LastName    string `json: "lastName"`
	Email       string `json: "email"`
	PhoneNumber string `json: "phoneNumber"`
}

type RegistrationResponse struct {
	Data  string `json: "Registration Complete"`
	Error string `json: "error, omitempty"`
}
