package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewDeleteAccount(s *sorm.DeleteAccountRequest) *models.DeleteAccount {
	return &models.DeleteAccount{
		EventType:     "delete_account",
		CorrelationId: s.CorrelationId,
		TelcoId:       s.TelcoId,
		UserType:      s.UserType,
		AbInfo: models.DeleteAccountAbInfo{
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
			//ServiceId:       s.ServiceId,
			ImId:          ToImIdArray(s.ImIds),
			Additional:    ToAdditionalArray(s.Additional),
			DatetimeUnreg: s.DatetimeUnreg,
			ContractDate:  s.ContractDate,
		},
	}
}
