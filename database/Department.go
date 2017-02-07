package database

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name string
	Head Employee
	HeadID uint
	Employees []Employee
}
