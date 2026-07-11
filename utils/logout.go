package utils

import (
	"authenticatiion-flow/database"
)

func Logout() bool {
	getActived := database.GetActived()
	actived := *getActived
	var indexActived int
	for i := range actived {
		if actived[i].Status == true {
			indexActived = i
		}
	}
	actived[indexActived].Status = false
	return true
}
