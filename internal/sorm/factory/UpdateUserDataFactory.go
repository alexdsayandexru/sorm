package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func ToUpdateUserDataAbInfoAdditionalArray(adds []*sorm.Add) []models.UpdateUserDataAbInfoAdditional {
	var result []models.UpdateUserDataAbInfoAdditional
	for _, m := range adds {
		result = append(result, models.UpdateUserDataAbInfoAdditional{
			Title:   m.Title,
			Content: m.Content,
		})
	}
	return result
}

func ToUpdateUserDataAbInfoImIdArray(adds []*sorm.ImId) []models.UpdateUserDataAbInfoImId {
	var result []models.UpdateUserDataAbInfoImId
	for _, m := range adds {
		result = append(result, models.UpdateUserDataAbInfoImId{
			ServiceId:   m.ServiceId,
			ServiceName: m.ServiceName,
		})
	}
	return result
}

func NewUpdateUserData(s *sorm.UpdateUserDataRequest) *models.UpdateUserData {
	return &models.UpdateUserData{
		EventType:     "user_update_data",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: models.UpdateUserDataAbInfo{
			UserId:          s.UserId,
			Nick:            s.Nick,
			BirthDate:       s.BirthDate,
			Name:            s.Name,
			Family:          s.Family,
			Initial:         s.Initial,
			Msisdns:         ToMsisdnArray(s.Msisdns),
			Emails:          ToEmailArray(s.Emails),
			Address:         s.Address,
			DatetimeReg:     s.DatetimeReg,
			DatetimeUpdated: s.DatetimeUpdated,
			ServiceId:       s.ServiceId,
			ImId:            ToUpdateUserDataAbInfoImIdArray(s.ImIds),
			Additional:      ToUpdateUserDataAbInfoAdditionalArray(s.Additional),
			ContractDate:    s.ContractDate,
		},
		EventData: models.UpdateUserDataEventData{
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
