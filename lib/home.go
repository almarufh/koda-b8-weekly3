package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"os"
	"regexp"
)

func Menu() {
	users := database.GetUsers()
	fmt.Println(*users)
	var input string

	for {
		utils.Clear()
		fmt.Printf("\n")
		fmt.Printf("--- Welcome to %s ---\n", database.Project)
		fmt.Printf("\n1. Register\n2. Forgot Password\n\n0. Exit\n\n")
		fmt.Printf("Note :")
		fmt.Printf("\n    - Login use username your account\n    - No have account must register first chose 1\n    - If you forget your password account you can chose 2\n")
		fmt.Printf("\nUername        : ")
		fmt.Scanf("%s", &input)

		value, _ := regexp.MatchString("^[a-zA-Z]{4,}$", input)

		if !value {
			switch input {
			case "1":
				Register(users)
				continue
			case "2":
				ChangePassword()
				continue
			case "0":
				os.Exit(0)
			default:
				utils.WrongInput()
				continue
			}
		}

		Login(input)
	}
}
