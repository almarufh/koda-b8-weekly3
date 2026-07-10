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
	var input int

	for {
		utils.Clear()
		fmt.Println("\n")
		fmt.Printf("--- Welcome to %s ---\n", database.Project)
		fmt.Printf("\n1. Show Menu Product \n2. Register\n3. Login\n4. Forgot Password\n\n0. Exit\n\n")
		fmt.Printf("Choose a menu : ")

		_, err := fmt.Scanln(&input)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch input {
		case 1:
			utils.Clear()
			ShowCategory()
		case 2:
			Register(users)
			return
		case 3:
			Login()
			return
		case 4:
			ChangePassword()
			return
		case 0:
			os.Exit(0)
		default:
			continue
		}
	}
}
