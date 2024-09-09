package response

type ApiResponse struct {
	Success bool        `json: "success"`
	Data    interface{} `json: "data, omitempty"`
}
