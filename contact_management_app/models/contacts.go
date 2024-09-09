package models

type ContactDetails struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

func (o *ContactDetails) GetFirstName() string {
	return o.FirstName
}

func (o *ContactDetails) SetFirstName(firstName string) {
	o.FirstName = firstName
}

func (o *ContactDetails) GetLastName() string {
	return o.LastName
}

func (o *ContactDetails) SetLastName(lastName string) {
	o.LastName = lastName
}

func (o *ContactDetails) GetEmail() string {
	return o.Email
}

func (o *ContactDetails) SetEmail(email string) {
	o.Email = email
}

func (o *ContactDetails) GetPhoneNumber() string {
	return o.PhoneNumber
}

func (o *ContactDetails) SetPhoneNumber(phoneNumber string) {
	o.PhoneNumber = phoneNumber
}
