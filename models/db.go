package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB       *gorm.DB
	host     = os.Getenv("HOST")
	port     = os.Getenv("PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
	dbInfo   = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func SetupModels() *gorm.DB {
	DB, err := gorm.Open("postgres", "postgres://mdqtemsu:gvtDbsKCdF6eUiD6Ap0nahB7pJSzZ0wA@john.db.elephantsql.com:5432/mdqtemsu") //dbInfo) // Or use this
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&History{})

	return DB
}
