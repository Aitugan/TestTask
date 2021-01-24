package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Aitugan/currapi/models"
	"github.com/Aitugan/currapi/services/routes"
	"github.com/Aitugan/currapi/utils"
)

func init() {
	models.DB = models.SetupModels()
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(1).Day().Do(utils.UpdateCurrencies)
	s1.StartAsync()
}

func main() {
	router := routes.NewRoutes()

	fmt.Println("Connected to 8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}
