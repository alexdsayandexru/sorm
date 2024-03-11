package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewUserAccountRecovery(s *sorm.UserAccountRecoveryRequest) *models.UserAccountRecovery {
	return &models.UserAccountRecovery{
		EventType:     "user_account_recovery",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		EventData: models.UserAccountRecoveryEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Msisdn:    s.MsisdnLogin,
			Email:     s.EmailLogin,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
		},
	}
}
