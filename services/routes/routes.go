package routes

import (
	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Aitugan/currapi/handlers"
)

func NewRoutes() *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api").Subrouter()

	base.HandleFunc("/exchange-rate", handlers.GetExchangeRate).Methods("GET")
	base.HandleFunc("/convert", handlers.Convert).Methods("POST")
	base.HandleFunc("/history", handlers.GetRateHistory).Methods("GET")

	return router
}
