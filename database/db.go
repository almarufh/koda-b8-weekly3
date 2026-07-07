package database

// Variable User
type Users struct {
	first    string
	last     string
	email    string
	password string
}

var users []Users = []Users{}

// Variable User Actived
type UserActive struct {
	name     string
	email    string
	password string
	status   bool
}

var actived []UserActive = []UserActive{}

// Func Output Memory Adress
func allVariable() (*[]Users, *[]UserActive) {
	return &users, &actived
}
