package database

import "github.com/jinzhu/gorm"

type Employee struct{
	gorm.Model
	Name string
	DepartmentID uint
	TaskID uint
}
