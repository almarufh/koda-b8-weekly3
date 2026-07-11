package database

// Variable User
type Users struct {
	First    string
	Last     string
	Email    string
	Password string
}

// Variable User Actived
type UserActive struct {
	Name     string
	Email    string
	Password string
	Status   bool
}

// Type data results func
type Results struct {
	Message string
	Code    string
	Data    Users
}

type Products struct {
	Id          string
	Name        string
	Category    string
	Description string
	Price       int
	IsReady     bool
	IsPromo     bool
	Stock       int
}

type Category struct {
	Urut        int
	Id          string
	Name        string
	Description string
	Price       int
	Stock       int
}

var cart []string

var productsCategory []Category = []Category{}

var Project string = "KFC Depok Sawangan"

var products []Products = []Products{}

var actived []UserActive = []UserActive{}

var users []Users = []Users{
	{
		First:    "Alma'ruf",
		Last:     "Hidayat",
		Email:    "owner@gmail.com",
		Password: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	{
		First:    "Alma'ruf",
		Last:     "Hidayat",
		Email:    "owner",
		Password: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	{
		First:    "Ali",
		Last:     "Ghufro",
		Email:    "admin",
		Password: "81dc9bdb52d04dc20036dbd8313ed055",
	},
	{
		First:    "CuanBot",
		Last:     "",
		Email:    "a",
		Password: "0cc175b9c0f1b6a831c399e269772661",
	},
}

func GetCart() *[]string {
	return &cart
}

func GetProducts() *[]Products {
	return &products
}

func GetCategory() *[]Category {
	return &productsCategory
}

func GetActived() *[]UserActive {
	return &actived
}

func GetUsers() *[]Users {
	return &users
}

// Methode

// Func Get Email, Password, Full Name
func (u Users) FullName() string {
	return u.First + " " + u.Last
}

func (u Users) GetEmail() string {
	return u.Email
}

func (u Users) GetPassword() string {
	return u.Password
}
