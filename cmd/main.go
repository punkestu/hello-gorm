package main

import (
	"fmt"
	"github.com/punkestu/hello-gorm/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// create connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create record
	db.Create(&entity.Object{Component1: "hello gorm"})

	// get first record
	var o1 entity.Object
	db.First(&o1)
	fmt.Println("first data:", o1.Component1)
	var o2 entity.Object
	db.First(&o2, 1)
	fmt.Println("first data where id:", o2.Component1)
	// same with this db.First(&object, "id=?", id)

	// get all records
	var os []entity.Object
	db.Find(&os)
	fmt.Println("all data:")
	for _, o := range os {
		fmt.Println(o)
	}

	// update record
	db.Model(&o1).Update("component1", "hello updated")
	fmt.Println("updated data:", o1.Component1)
	// same with this
	// db.Model(&o1).Updates(entity.Object{Component1: "hello updated"})

	// delete record
	db.Delete(&o1)
	fmt.Println("deleted data:", o1.Component1)
}
