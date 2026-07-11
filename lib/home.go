package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
)

func Menu() {
	users := database.GetUsers()
	fmt.Println(*users)
	var input string

	for {
		utils.Clear()
		fmt.Printf("\n")
		fmt.Printf("--- Welcome to %s ---\n", database.Project)
		fmt.Printf("\n1. Register\n2. Login\n3. Forgot Password\n\n0. Exit\n\n")
		fmt.Printf("Choose a menu : ")
		fmt.Scanf("%s", &input)
		switch input {
		case "1":
			Register(users)
			return
		case "2":
			Login()
			return
		case "4":
			ChangePassword()
			return
		case "0":
			os.Exit(0)
		default:
			utils.WrongInput()
			continue
		}
	}
}
