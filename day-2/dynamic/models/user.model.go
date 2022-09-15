package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int    `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
