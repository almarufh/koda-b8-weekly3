package utils

import "fmt"

func WrongInput() {
	fmt.Printf("\nAlert : Input failed !!!\n")
	fmt.Printf("\nPress ENTER for try again...")
	fmt.Scanln()
}
