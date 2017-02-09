package database

type Task struct {
	BaseEntity
	Description string
	Responsible []Employee
}
