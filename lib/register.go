package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
)

func Register(users *[]database.Users) {
	var first string
	var last string
	var email string

	for {
		utils.Clear()
		fmt.Printf("\n\n--- REGISTER ---\n\nFirst Name :  ")
		fmt.Scanf("%s", &first)

		fmt.Printf("Last Name :  ")
		fmt.Scanf("%s", &last)

		fmt.Printf("Username :  ")
		fmt.Scanf("%s", &email)

		password := CreatePassword()

		newUser := database.Users{
			First:    first,
			Last:     last,
			Email:    email,
			Password: utils.Encrypted(password),
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("\nEmail belum terdaftar : %v\n", r)
			}
		}()

		utils.Clear()
		var confirmAccept string
		fmt.Printf("\n\nIs it true?\n\n")
		fmt.Println("First Name : ", first)
		fmt.Println("Last Name : ", last)
		fmt.Println("Username : ", email)
		fmt.Printf("\n\nContinue (y/n) :  ")
		fmt.Scanf("%s", &confirmAccept)
		switch confirmAccept {
		case "y":
			*users = append(*users, newUser)
			fmt.Printf("\n\nRegister success, press enter to back..")
			fmt.Scanf("\n")
			return
		case "n":
			continue
		default:
			utils.WrongInput()
			continue
		}
	}
}
