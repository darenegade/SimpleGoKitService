package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	database *gorm.DB
)

func Initialize() {

	var err error

	database, err = gorm.Open("mysql", "root:secret@tcp(mysql:3306)/tasksgo?charset=utf8&parseTime=True&loc=Local")
	//database, err = gorm.Open("mysql", "root:secret@/tasksgo?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	database.CreateTable(&Employee{}, &Task{}, &Department{})
}

type Repository interface {
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(ID uint) error
	FindOne(entity interface{}) error
	FindAll(entities interface{}) error
}

func Create(entity interface{}) error {

	if database.NewRecord(entity) {

		database.Create(entity)
		return nil
	}

	return errors.New("Entity already exists.")
}

func Update(entity interface{}) error {

	if !database.NewRecord(entity) {

		database.Save(entity)
		return nil
	}

	return errors.New("Entity doesn't exists.")
}

func FindOne(entity interface{}, ID uint) error {

	database.Find(entity, ID)

	if !database.NewRecord(entity) {
		return nil
	} else {
		return errors.New("Entity doesn't exists.")
	}

}

func FindAll(entities interface{}) error {

	database.Find(entities)
	return nil
}

func Delete(entity interface{}) error {

	database.Delete(entity)
	return nil
}
