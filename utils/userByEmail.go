package utils

import (
	"authenticatiion-flow/database"
	"strings"
)

type UserByEmail struct {
	Email    string
	Password string
	Name     string
}

type ResultByEmail struct {
	Status bool
	Index  int
	Data   UserByEmail
}

func GetUserByEmail(email string) ResultByEmail {
	users := database.GetUsers()

	for i, res := range *users {
		emailLower := strings.ToLower(res.Email)
		searchEmail := strings.ToLower(email)
		if emailLower == searchEmail {
			result := ResultByEmail{
				Status: true,
				Index:  i,
				Data: UserByEmail{
					Email:    res.GetEmail(),
					Password: res.GetPassword(),
					Name:     res.FullName(),
				},
			}
			return result
		}
	}

	return ResultByEmail{
		Status: false,
		Index:  -1,
		Data:   UserByEmail{},
	}
}
