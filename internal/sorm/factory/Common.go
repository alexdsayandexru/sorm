package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func ToMsisdnArray(msisdns []*sorm.Msisdn) []string {
	var result []string
	for _, m := range msisdns {
		result = append(result, m.Msisdn)
	}
	return result
}

func ToEmailArray(emails []*sorm.Email) []string {
	var result []string
	for _, m := range emails {
		result = append(result, m.Email)
	}
	return result
}

func ToAdditionalArray(adds []*sorm.Add) []models.Additional {
	var result []models.Additional
	for _, m := range adds {
		result = append(result, models.Additional{
			Title:   m.Title,
			Content: m.Content,
		})
	}
	return result
}

func ToImIdArray(adds []*sorm.ImId) []models.ImId {
	var result []models.ImId
	for _, m := range adds {
		result = append(result, models.ImId{
			ServiceId:   m.ServiceId,
			ServiceName: m.ServiceName,
		})
	}
	return result
}
