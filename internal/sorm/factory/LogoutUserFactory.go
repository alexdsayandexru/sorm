package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewLogoutUser(s *sorm.LogoutUserRequest) *models.LogoutUser {
	return &models.LogoutUser{
		EventType:     "user_login",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		Event: models.LogoutUserEvent{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Ip:        s.Ip,
			Port:      s.Port,
			UserAgent: s.UserAgent,
			Datetime:  s.Datetime,
		},
	}
}
