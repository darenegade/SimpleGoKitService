package database

type Department struct {
	BaseEntity
	Name string
	// Has One Head of Dep.
	Head   Employee `gorm:"ForeignKey:HeadID;AssociationForeignKey:HeadID"`
	HeadID uint     `json:"-"`
	// Has Many Employees
	EmployeesID uint       `json:"-"`
	Employees   []Employee `gorm:"ForeignKey:DepartmentID;AssociationForeignKey:EmployeesID"`
}
