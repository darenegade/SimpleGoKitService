package database

type Employee struct {
	BaseEntity
	Name         string
	DepartmentID uint `json:"-"`
	HeadID uint `json:"-"`
}
