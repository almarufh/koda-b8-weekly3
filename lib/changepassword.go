package lib

import (
	"authenticatiion-flow/database"
	"authenticatiion-flow/utils"
	"fmt"
	"strings"
)

func ChangePassword() {
	var email string
	getUsers := database.GetUsers()
	search := utils.ResultByEmail{}

first:
	for {
		utils.Clear()
		fmt.Printf(`Input "Back" and press enter to back`)
		fmt.Printf("\n\n")
		fmt.Printf("\n--- [Forgot Password ] ---\n\nInput your username :  ")
		fmt.Scanf("%s", &email)
		if strings.ToLower(email) == "back" {
			return
		}

		search = utils.GetUserByEmail(email)
		found := search.Status

		if !found {
			fmt.Printf("\nUsername not found !!!\n\nPress ENTER to try again")
			fmt.Scanf("\n")
			continue
		}

		break
	}

	for {
		utils.Clear()
		fmt.Printf("--- Whats This Your Account ---\n\n")
		fmt.Printf("Name      : %s\n", search.Data.Name)
		fmt.Printf("Username  : %s\n", search.Data.Email)

		fmt.Printf("\n\nConfirm (y/n) : ")
		fmt.Scanf("%s", &email)

		switch strings.ToLower(email) {
		case "y":
			newPassword := CreatePassword(search.Data.Name, search.Data.Email)

			(*getUsers)[search.Index].Password = utils.Encrypted(newPassword)

			utils.Clear()
			fmt.Printf("Password changed successfully!\n\nPress enter to home menu... ")
			fmt.Scanf("\n")
			return
		case "n":
			goto first
		default:
			continue
		}
	}
}
