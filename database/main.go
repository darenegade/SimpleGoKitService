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
	database, err = gorm.Open("mysql", "root:@localhost:3306/tasksgo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Employee{}, &Task{}, &Department{})
}

type Repository interface {
	create(entity interface{}) error
	update(entity interface{}) error
	delete(ID uint) error
	findOne(entity interface{}, ID uint) error
	findAll(entities []interface{}) error
}

type SQLRepository struct{}

func (SQLRepository) create(entity interface{}) error {

	if database.NewRecord(entity) {

		database.Create(&entity)
		return nil
	}

	return errors.New("Entity already exists.")
}

func (SQLRepository) update(entity interface{}) error {

	if !database.NewRecord(entity) {

		database.Update(&entity)
		return nil
	}

	return errors.New("Entity doesn't exists.")
}

func (SQLRepository) findOne(entity interface{}, ID uint) error {

	if !database.NewRecord(entity) {

		database.Find(&entity, ID)
		return nil
	}

	return errors.New("Entity doesn't exists.")
}

func (SQLRepository) findAll(entities []interface{}, ID uint) error {

	database.Find(&entities)
	return nil
}

func (SQLRepository) delete(entity interface{}) error {

	database.Delete(&entity)
	return nil
}
