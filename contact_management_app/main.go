package main

import (
	"fmt"
	"net/http"
	"project/web"
)

func main() {

	http.HandleFunc("/register", web.RegisterUser)
	http.ListenAndServe(":8080", nil)
	fmt.Print("Server is running")
}
