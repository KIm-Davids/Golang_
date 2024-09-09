package request

type LoginRequest struct {
	UserName string `json: "username"`
	Password string `json: "password"`
}

type LoginResponse struct {
	Data  string `json: "User Logged in Successfully"`
	Error string `json: "error, omitempty"`
}
