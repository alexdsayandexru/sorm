package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewDeleteUserAccount(s *sorm.DeleteUserAccountRequest) *models.DeleteUserAccount {
	return &models.DeleteUserAccount{
		EventType:     "delete_user_account",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		EventData: models.DeleteUserAccountEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
			Message:   s.Message,
		},
	}
}
