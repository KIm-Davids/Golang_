package main

import (
	"fmt"
	"net/http"
	"project/repository"
	_ "project/repository"
	"project/web"
)

func main() {

	repository.ConnectToMySql()
	repository.CreateTableToInsertContacts()
	http.HandleFunc("/register-user", web.RegisterUser)
	http.HandleFunc("/login", web.LoginUser)
	http.HandleFunc("/get-all-contacts", web.GetAllContacts)
	http.HandleFunc("/create-contact", web.CreateContact)
	fmt.Print("Server is running")
	http.ListenAndServe("localhost:8080", nil)
}
