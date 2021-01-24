package handlers

import (
	"net/http"

	"github.com/Aitugan/currapi/services"
	"github.com/Aitugan/currapi/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetRateHistory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	data, resp := services.GetRateHistoryService(query)
	if resp != "Success" {
		respond := utils.Message(false, resp)
		utils.Respond(w, respond)
	}
	respond := utils.Message(true, resp)
	respond["data"] = data
	utils.Respond(w, respond)
}
