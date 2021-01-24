package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Aitugan/currapi/models"
	"github.com/Aitugan/currapi/services"
	"github.com/Aitugan/currapi/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	var conv models.Conversion
	json.NewDecoder(r.Body).Decode(&conv)

	data, resp := services.ConvertService(conv)
	if resp != "Success" {
		respond := utils.Message(false, resp)
		utils.Respond(w, respond)
	}
	respond := utils.Message(true, resp)
	respond["data"] = fmt.Sprintf("%.2f", data)
	utils.Respond(w, respond)
}
