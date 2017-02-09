package database

type Department struct {
	BaseEntity
	Name      string
	Head      Employee
	HeadID    uint `json:"-"`
	Employees []Employee
}
