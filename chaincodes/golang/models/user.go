package models

type User struct {
	UserId    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	ContactNo string `json:"contactNo"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}

type Participant struct {
	User   User   `json:"user"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
