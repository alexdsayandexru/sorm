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

func NewRegisterUser(s *sorm.RegisterUserRequest) *models.RegisterUser {
	return &models.RegisterUser{
		EventType:     "user_registration",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: models.AbInfo{
			UserId:       s.UserId,
			Msisdns:      ToMsisdnArray(s.Msisdns),
			Emails:       ToEmailArray(s.Emails),
			DatetimeReg:  s.DatetimeReg,
			ServiceId:    s.ServiceId,
			Additional:   ToAdditionalArray(s.Additional),
			ContractDate: s.ContractDate,
		},
		EventData: models.RegisterUserEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Msisdn:    s.MsisdnLogin,
			Email:     s.EmailLogin,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
			Message:   s.Message,
		},
	}
}
