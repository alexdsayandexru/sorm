package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewLoginUser(s *sorm.LoginUserRequest) *models.LoginUser {
	return &models.LoginUser{
		EventType:     "user_login",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		Event: models.LoginUserEvent{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Msisdn:    s.MsisdnLogin,
			Email:     s.EmailLogin,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Factor:    s.Factor,
			Datetime:  s.Datetime,
		},
	}
}
