package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

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
