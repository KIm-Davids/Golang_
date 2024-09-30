package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	request2 "project/dtos/request"
	"project/models"
)

func ConnectToMySql() {
	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	if err != nil {
		log.Fatal(err)
	}

	CreateTable(db)

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to the database successfully")
}

func CreateTable(db *sql.DB) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
 		last_name VARCHAR(100) NOT NULL,
 		email VARCHAR(250) NOT NULL,
 		phone_number VARCHAR(20) NOT NULL,
		password VARCHAR(100) NOT NULL,
		login_status VARCHAR(20) NOT NULL
	)`

	_, err := db.Exec(createTableSQL)

	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Table created successfully")

}

func InsertUsersIntoDatabase(firstName string, lastName string, email string, phoneNumber string, password string) {
	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()

	createQuery := `INSERT INTO users (
    first_name,
    last_name,
    email,
    phone_number,
    password,
    login_status               
) VALUES (?, ?, ?, ?, ?, ?)`

	loginStatus := "offline"
	_, err = db.Exec(createQuery, firstName, lastName, email, phoneNumber, password, loginStatus)

	if err != nil {
		log.Fatalf("An error occurred %v", err)
	}
}

func RetrieveValuesFromTheDb() []models.UserDetails {
	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	query := `
	SELECT email, password, login_status
	FROM golangPhoneBook.users
	WHERE login_status = 'offline'
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error retrieving values from db %v", err)
	}

	defer rows.Close()

	var users []models.UserDetails

	for rows.Next() {
		var user models.UserDetails
		if err := rows.Scan(&user.Email, &user.Password, &user.LoginStatus); err != nil {
			log.Fatalf("Error occurred during iteration %v", err)
		}
		users = append(users, user)
	}
	if err != nil {
		log.Fatalf("An error just occurred during the retrieval of users %v", err)
	}
	return users
}

func SetLoginStatus(loginStatus string, email string, password string) {

	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}

	defer db.Close()

	query := `
		UPDATE users
		SET login_status = ?
		WHERE login_status = 'offline' AND email = ? AND password = ?
	`

	_, err = db.Exec(query, loginStatus, email, password)

	if err != nil {
		log.Fatalf("Could not set login status %v", err)
	}
}

//func RetrievePassword() interface{} {
//
//	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")
//
//	if err != nil {
//		log.Fatalf("Failed to log connect to database %v", err)
//	}
//
//	query := `
//		SELECT password
//		FROM users
//	`
//
//	row, err := db.Query(query)
//
//	for row.Next() {
//		var user models.UserDetails
//
//		err := row.Scan(&user.Password)
//
//		password := user
//		if err != nil {
//			log.Fatalf("Error in retrieving data %v", err)
//		}
//		return password
//	}
//	return nil
//}

func CreateTableToInsertContacts() {
	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	if err != nil {
		log.Fatalf("Couldn't connect to the database %v", err)
	}
	query := `
		CREATE TABLE IF NOT EXISTS golang_phone_book (
		contact_id INT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR (100) NOT NULL,
		last_name VARCHAR (100) NOT NULL,
		phone_number VARCHAR (20) UNIQUE NOT NULL,
		email VARCHAR (250) UNIQUE NOT NULL,
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE 
	)`

	_, err = db.Exec(query)

	if err != nil {
		log.Fatalf("An error occurred in connecting to database %v", err)
	}

}

func GetLoginStatusOfUser() models.UserDetails {

	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	query := `
		SELECT login_status
		FROM users
		WHERE login_status = 'online'
	`

	row, err := db.Query(query)

	var user models.UserDetails

	for row.Next() {
		err = row.Scan(&user.LoginStatus)
	}

	if err != nil {
		log.Fatalf("Error retrieving values from the db %v", err)
	}

	loginStatus := user

	return loginStatus
}

func SaveContactToTheDb(request request2.CreateContactRequest) bool {
	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	var user models.UserDetails

	query := `
		INSERT INTO golang_phone_book( 
 		first_name,
		last_name,
		email,
		phone_number,
		user_id 
		)VALUES (?, ?, ?, ?, ?)
		
	`

	retrieveEmailQuery := `
		SELECT id FROM users
		WHERE login_status = ?
	`

	login := "online"

	row, err := db.Query(retrieveEmailQuery, login)

	for row.Next() {
		err = row.Scan(&user.Id)
	}

	_, err = db.Exec(query, request.FirstName, request.LastName, request.Email, request.PhoneNumber, user.Id)

	if err != nil {
		log.Fatalf("Error saving contact to the database %v", err)
		return false
	}
	return true
}

func RetrieveAllValuesForUser() ([]models.ContactDetails, error) {

	var user models.UserDetails

	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	retrieveContactQuery := `
	SELECT first_name, last_name, email, phone_number
	FROM golang_phone_book
	WHERE user_id = ?
	`

	retrieveUserId := `
		SELECT id FROM users
		WHERE login_status = 'online'
	`

	userId, err := db.Query(retrieveUserId)

	fmt.Print(userId)

	for userId.Next() {
		err = userId.Scan(&user.Id)

	}

	contacts, err := db.Query(retrieveContactQuery, user.Id)

	//fmt.Print(contacts)
	var contactList []models.ContactDetails

	var contact models.ContactDetails
	var usersContactList []models.ContactDetails

	for contacts.Next() {
		contacts.Scan(&contact.FirstName, &contact.LastName, &contact.Email, &contact.PhoneNumber)
		usersContactList = append(contactList, contact)
	}

	if err != nil {
		log.Fatal("Error in retrieving all contacts ")
	}

	return usersContactList, err
}

func RetrieveContactFromContactDbForValidation() map[models.ContactDetails]string {
	db, err := sql.Open("mysql", "root:Kimdavid1.@tcp(127.0.0.1:3306)/golangPhoneBook")

	retrieveContactQuery := `
		SELECT g.email, g.phone_number
		FROM golang_phone_book g 
		JOIN users u ON g.user_id = u.id 
	`

	contacts, err := db.Query(retrieveContactQuery)

	if err != nil {
		log.Fatalf("Eror retrieving values from the database %v", err)
	}

	contactList := make(map[models.ContactDetails]string)
	var contactDetails models.ContactDetails

	for contacts.Next() {
		err = contacts.Scan(&contactDetails.PhoneNumber, &contactDetails.Email)
		
	}
	//fmt.Print(contactDetails)
	if err != nil {
		log.Fatalf("Eror retrieving values from the database %v", err)
	}

	return contactList

}
