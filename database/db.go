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

// Func Output Memory Adress
func allVariable() (*[]Users, *[]UserActive) {
	return &users, &actived
}
