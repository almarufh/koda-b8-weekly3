package lib

import (
	"authenticatiion-flow/utils"
	"fmt"
	"regexp"
	"strings"
)

func Search() {
	for {
		var search string
		utils.Clear()
		fmt.Printf(`Input "Back" and press ENTER for back`)
		fmt.Printf("\n\n")
		fmt.Printf("\nSearch what you need products :  ")
		fmt.Scanln(&search)

		if strings.ToLower(search) == "back" {
			return
		}

		value, _ := regexp.MatchString("^[a-zA-Z]{3,}$", search)

		if !value {
			utils.Clear()
			fmt.Printf("Keyword must min 4 character !!!\n\nPress ENTER to retry... ")
			fmt.Scanln()
			continue
		}

		SearchProduct(search)
	}
}
