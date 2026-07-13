package utils

import (
	"authenticatiion-flow/database"
)

func NameActived() string {
	getActived := database.GetActived()
	actived := *getActived
	var name string
	for i := range actived {
		if actived[i].Status == true {
			name = actived[i].Name
		}
	}
	return name
}

func UserNameActived() string {
	getActived := database.GetActived()
	actived := *getActived
	var name string
	for i := range actived {
		if actived[i].Status == true {
			name = actived[i].Email
		}
	}
	return name
}
