package database

type Employee struct {
	BaseEntity
	Name         string
	DepartmentID uint
	TaskID       uint
}
