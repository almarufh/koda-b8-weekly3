package utils

import "fmt"

func WrongInput(input *string) {
	fmt.Println("\nInput failed ...")
	fmt.Printf("\nPress ENTER to chose again...")
	fmt.Scanln(input)
}
