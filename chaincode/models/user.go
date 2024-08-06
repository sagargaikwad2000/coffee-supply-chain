package models

type User struct {
	UserId    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email,omitempty"`
	ContactNo string `json:"contactNo"`
	Address   string `json:"address"`
}

type Participant struct {
	User
	Role   string `json:"role"`
	Status string `json:"status"`
}
