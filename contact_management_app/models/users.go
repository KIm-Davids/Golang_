package models

type UserDetails struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	LoginStatus string
	Id          int
}

func (o *UserDetails) GetFirstName() string {
	return o.FirstName
}

func (o *UserDetails) SetFirstName(firstName string) {
	o.FirstName = firstName
}

func (o *UserDetails) GetLastName() string {
	return o.LastName
}

func (o *UserDetails) SetLastName(lastName string) {
	o.LastName = lastName
}

func (o *UserDetails) GetEmail() string {
	return o.Email
}

func (o *UserDetails) SetEmail(email string) {
	o.Email = email
}

func (o *UserDetails) GetPhoneNumber() string {
	return o.PhoneNumber
}

func (o *UserDetails) SetPhoneNumber(phoneNumber string) {
	o.PhoneNumber = phoneNumber
}
