package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/GOLANG_RESTAPI"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migrate the products table
	err = database.AutoMigrate(&Product{})
	if err != nil {
		panic("Failed to migrate the database!")
	}

	DB = database
}
