package database

import "github.com/jinzhu/gorm"

type Task struct{
	gorm.Model
	Description string
	Responsible []Employee
}