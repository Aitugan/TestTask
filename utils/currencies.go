package utils

import (
	"strconv"

	"github.com/Aitugan/currapi/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var URL = "https://nationalbank.kz/rss/rates_all.xml"

func UpdateCurrencies() {
	ReadURL(URL)
	for _, v := range AllData.Chan.Curr {
		h := &models.History{}
		h.Title = v.Title
		h.PubDate = v.PubDate
		currDescr, err := strconv.ParseFloat(v.Description, 64)
		Check(err)
		currChange, err := strconv.ParseFloat(v.Change, 64)
		Check(err)
		h.Description = currDescr - currChange
		h.Change, err = strconv.ParseFloat(v.Change, 64)
		Check(err)
		models.DB.Create(&h)
	}
}

var AllData models.RSS

func FindItemByCurrency(currency string) *models.Currency {
	curren := models.Currency{}
	for _, v := range AllData.Chan.Curr {
		if v.Title == currency {
			curren = v
			return &curren
		}
	}
	return nil
}
