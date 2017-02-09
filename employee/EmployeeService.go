package employee

import (
	"github.com/darenegade/SimpleGoKitService/database"
	"github.com/darenegade/SimpleGoKitService/util"
)



type EmployeeRepository interface {
	findAll() ([]database.Employee, error)
	findOne(uint) (database.Employee, error)
	create(database.Employee) (database.Employee, error)
	update(database.Employee) (database.Employee, error)
	delete(uint) error
}

type EmployeeService struct {}

func (EmployeeService) findAll() ([]database.Employee,error) {
	var employees []database.Employee
	err := database.FindAll(&employees)

	return employees, err
}

func (EmployeeService) findOne(ID uint) (database.Employee, error){
	var employee database.Employee
	err := database.FindOne(&employee, ID)

	return employee, err
}

func (EmployeeService) create(employee database.Employee) (database.Employee, error){
	err := database.Create(&employee)

	return employee, err
}

func (es EmployeeService) update(employee database.Employee, ID uint) (database.Employee, error){

	current, err := es.findOne(ID)

	if err != nil {
		return employee, err
	}

	current.Name = employee.Name

	err = database.Update(&current)

	return current, err
}

func (EmployeeService) delete(ID uint) error{
	var employee database.Employee
	employee.ID = ID
	err := database.Delete(&employee)

	return err
}