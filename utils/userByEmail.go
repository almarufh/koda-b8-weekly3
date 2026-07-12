package utils

import (
	"authenticatiion-flow/database"
	"strings"
)

type UserByEmail struct {
	email    string
	password string
	name     string
}

type ResultByEmail struct {
	status bool
	data   UserByEmail
}

func GetUserByEmail(email string) ResultByEmail {
	users := database.GetUsers()

	for _, res := range *users {
		emailLower := strings.ToLower(res.Email)
		searchEmail := strings.ToLower(email)
		if emailLower == searchEmail {
			result := ResultByEmail{
				status: true,
				data: UserByEmail{
					email:    res.GetEmail(),
					password: res.GetPassword(),
					name:     res.FullName(),
				},
			}
			return result
		}
	}

	return ResultByEmail{
		status: false,
		data:   UserByEmail{},
	}
}
