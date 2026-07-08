package database

// Variable User
type Users struct {
	First    string
	Last     string
	Email    string
	Password string
}

var users []Users = []Users{}

// Variable User Actived
type UserActive struct {
	Name     string
	Email    string
	Password string
	Status   bool
}

var actived []UserActive = []UserActive{}

// Type data results func
type Results struct {
	Message string
	Code    string
	Data    Users
}

// Func Output Memory Adress
func AllVariable() (*[]Users, *[]UserActive) {
	return &users, &actived
}
