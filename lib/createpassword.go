package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
	"regexp"
)

func CreatePassword(name string, username string) string {
	var password1 string
	var password2 string

	for {
		utils.Clear()
		fmt.Printf("Full Name : %s\n", name)
		fmt.Printf("Username  : %s\n", username)
		fmt.Printf("Create New PIN (4 digits):  ")
		fmt.Scanln(&password1)

		value, _ := regexp.MatchString("^[0-9]{4}$", password1)

		if !value {
			utils.Clear()
			fmt.Printf("PIN must be exactly 4 digits of numbers!!!\n\nPress ENTER to retry... ")
			fmt.Scanln()
			continue
		}

		fmt.Printf("Confirm New PIN:  ")
		fmt.Scanln(&password2)

		if password1 == password2 {
			return password1
		} else {
			utils.Clear()
			fmt.Printf("Wrong confirm password!!!\n\nPress ENTER to retry! ")
			fmt.Scanln()
			continue
		}
	}
}
