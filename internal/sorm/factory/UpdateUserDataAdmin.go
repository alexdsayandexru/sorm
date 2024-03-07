package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewUpdateUserDataAdmin(s *sorm.UpdateUserDataAdminRequest) *models.UpdateUserDataAdmin {
	return &models.UpdateUserDataAdmin{
		EventType:     "user_update_admin",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: models.UpdateUserDataAdminAbInfo{
			UserId:          s.UserId,
			Msisdns:         ToMsisdnArray(s.Msisdns),
			Emails:          ToEmailArray(s.Emails),
			Nick:            s.Nick,
			BirthDate:       s.BirthDate,
			Name:            s.Name,
			Family:          s.Family,
			Initial:         s.Initial,
			Address:         s.Address,
			DatetimeReg:     s.DatetimeReg,
			DatetimeUpdated: s.DatetimeUpdated,
			ServiceId:       s.ServiceId,
			ImId:            ToImIdArray(s.ImIds),
			Additional:      ToAdditionalArray(s.Additional),
			ContractDate:    s.ContractDate,
		},
		EventData: models.UpdateUserDataAdminEventData{
			ServiceId: s.ServiceId,
			UserId:    s.UserId,
			Datetime:  s.Datetime,
		},
	}
}
