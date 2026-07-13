package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
)

func Login(inputEmail string) {
	var inputPassword string
	fmt.Printf("PIN (4 digits) : ")
	fmt.Scanf("%s", &inputPassword)
	fmt.Printf("\n--------------------\n")
	password := utils.Encrypted(inputPassword)
	CheckUser(inputEmail, password)
}
