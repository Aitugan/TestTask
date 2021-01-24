package services

import (
	"fmt"
	"net/url"

	"github.com/Aitugan/currapi/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetRateHistoryService(query url.Values) ([]*models.GetRHistoryResult, string) {
	base, baseExists := query["base"]
	if !baseExists || len(base) == 0 {
		return nil, "Base not defined"
	}
	quote, quoteExists := query["quote"]
	if !quoteExists || len(quote) == 0 {
		return nil, "Quote not defined"
	}

	var result []*models.GetRHistoryResult
	if quoteExists && quote[0] != "KZT" {
		var histBase []models.History
		if err := models.DB.Where("title = ?", base[0]).Find(&histBase).Error; err != nil {
			return nil, "Base not found"
		}
		var histQuote []models.History
		if err := models.DB.Where("title = ?", quote[0]).Find(&histQuote).Error; err != nil {
			return nil, "Quote not found"
		}
		for i := 0; i < len(histBase); i++ {
			rHist := &models.GetRHistoryResult{}
			rHist.Date = histBase[i].PubDate
			rHist.Value = fmt.Sprintf("%.2f", (histBase[i].Description-histBase[i].Change)/(histQuote[i].Description-histQuote[i].Change))
			result = append(result, rHist)
		}
		return result, "Success"
	} else {
		var hist []models.History
		if err := models.DB.Where("title = ?", base[0]).Find(&hist).Error; err != nil {
			return nil, "Base not found"
		}
		for _, v := range hist {
			rHist := &models.GetRHistoryResult{}
			rHist.Date = v.PubDate
			rHist.Value = fmt.Sprintf("%.2f", v.Description-v.Change)
			result = append(result, rHist)
		}
		return result, "Success"
	}
}
