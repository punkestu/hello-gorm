package main

import (
	"github.com/punkestu/hello-gorm/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Object{})
	if err != nil {
		panic(err)
	}
}
