package handlers

import (
	"net/http"

	"github.com/Aitugan/currapi/services"
	"github.com/Aitugan/currapi/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data, resp := services.GetExchangeRateService(query)
	if resp != "Success" {
		respond := utils.Message(false, resp)
		utils.Respond(w, respond)
	}
	respond := utils.Message(true, resp)
	respond["data"] = data
	utils.Respond(w, respond)
}
