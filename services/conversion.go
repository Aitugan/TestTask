package services

import (
	"strconv"

	"github.com/Aitugan/currapi/models"
	"github.com/Aitugan/currapi/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConvertService(conv models.Conversion) (float64, string) {
	var data float64
	if conv.ConvertFrom == "KZT" && conv.ConvertTo != "KZT" {
		itemQuote := utils.FindItemByCurrency(conv.ConvertTo)
		if itemQuote == nil {
			return 0, "Quote Doesn't exist"
		}
		quoteV, err := strconv.ParseFloat(itemQuote.Description, 64)
		utils.Check(err)
		data = conv.Amount / quoteV
		return data, "Success"

	} else if conv.ConvertFrom != "KZT" && conv.ConvertTo == "KZT" {
		itemBase := utils.FindItemByCurrency(conv.ConvertFrom)
		if itemBase == nil {
			return 0, "Base Doesn't exist"
		}
		baseV, err := strconv.ParseFloat(itemBase.Description, 64)
		utils.Check(err)

		data = conv.Amount * baseV
		return data, "Success"
	} else if conv.ConvertFrom == "KZT" && conv.ConvertTo == "KZT" {
		data = conv.Amount
		return data, "Success"
	} else {
		itemBase := utils.FindItemByCurrency(conv.ConvertFrom)
		if itemBase == nil {
			return 0, "Base Doesn't exist"
		}
		itemQuote := utils.FindItemByCurrency(conv.ConvertTo)
		if itemQuote == nil {
			return 0, "Quote Doesn't exist"
		}
		baseV, err := strconv.ParseFloat(itemBase.Description, 64)
		utils.Check(err)
		quoteV, err := strconv.ParseFloat(itemQuote.Description, 64)
		utils.Check(err)
		data = baseV * conv.Amount / quoteV
		return data, "Success"
	}
}
