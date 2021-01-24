package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func SetupModels() *gorm.DB {
	DB, err := gorm.Open("postgres", "postgres://mdqtemsu:gvtDbsKCdF6eUiD6Ap0nahB7pJSzZ0wA@john.db.elephantsql.com:5432/mdqtemsu")
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&History{})

	return DB
}
