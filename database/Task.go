package database

type Task struct {
	BaseEntity
	Description string
	// Many to Many - Responsible
	Responsible []Employee `gorm:"many2many:task_employees;"`
}
