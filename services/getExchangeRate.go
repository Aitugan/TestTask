package services

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Aitugan/currapi/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetExchangeRateService(query url.Values) (string, string) {
	base, baseExists := query["base"]
	if !baseExists || len(base) == 0 {
		return "", "Base not defined"
	}
	quote, quoteExists := query["quote"]
	var data string
	if quoteExists && quote[0] != "KZT" {
		itemBase := utils.FindItemByCurrency(base[0])
		if itemBase == nil {
			return "", "Base Doesn't exist"
		}
		itemQuote := utils.FindItemByCurrency(quote[0])
		if itemQuote == nil {
			return "", "Quote Doesn't exist"
		}
		baseV, err := strconv.ParseFloat(itemBase.Description, 64)
		utils.Check(err)
		quoteV, err := strconv.ParseFloat(itemQuote.Description, 64)
		utils.Check(err)
		data = fmt.Sprintf("%.2f", baseV/quoteV) + itemQuote.Title
	} else {
		item := utils.FindItemByCurrency(base[0])
		if item == nil {
			return "", "Base Doesn't exist"
		}
		data = item.Description + "KZT"
	}
	return data, "Success"
}
