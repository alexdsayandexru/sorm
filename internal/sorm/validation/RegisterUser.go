package validation

import (
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
	"github.com/alexdsayandexru/sorm/internal/validator"
)

func GetRegisterUserValidationRules(target *models.RegisterUser) ValidationRules {
	mapValidationRules := map[string]func() (bool, error){
		"correlation_id": func() (bool, error) {
			return validator.Validate(target.CorrelationId).Required().Uiid().GetResult()
		},
		"telco_id": func() (bool, error) {
			return validator.Validate(target.TelcoId).Required().Maximum(100).GetResult()
		},
		"user_type": func() (bool, error) {
			return validator.Validate(target.UserType).Required().Maximum(100).GetResult()
		},
		"user_id": func() (bool, error) {
			return validator.Validate(target.AbInfo.UserId).Required().MaxLength(255).Regex("^[A-Za-z0-9_-]+$").GetResult()
		},
		/*"msisdns": func() (bool, error) {
			return validator.Validate(target.Msisdns).RequiredIf(target.Emails == nil).MaxLength(16).Regex("^[0-9+]+$").GetResult()
		},
		"emails": func() (bool, error) {
			return validator.Validate(target.Emails).RequiredIf(target.Msisdns == nil).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
		},*/
		"datetime_reg": func() (bool, error) {
			return validator.Validate(target.AbInfo.DatetimeReg).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		/*"service_user": func() (bool, error) {
			return validator.Validate(target.ServiceUser).Required().Equal(1).GetResult()
		},*/
		"contract_date": func() (bool, error) {
			return validator.Validate(target.AbInfo.ContractDate).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
		"service_id": func() (bool, error) {
			return validator.Validate(target.AbInfo.ServiceId).Required().Maximum(100000000).GetResult()
		},
		"msisdn_login": func() (bool, error) {
			return validator.Validate(target.EventData.Msisdn).RequiredIf(len(target.EventData.Email) == 0).MaxLength(16).Regex("^[0-9+]+$").GetResult()
		},
		"email_login": func() (bool, error) {
			return validator.Validate(target.EventData.Email).RequiredIf(len(target.EventData.Msisdn) == 0).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
		},
		"ip": func() (bool, error) {
			return validator.Validate(target.EventData.Ip).MaxLength(255).Regex("^[0-9.]+$").GetResult()
		},
		"port": func() (bool, error) {
			return validator.Validate(target.EventData.Port).Maximum(99999).GetResult()
		},
		"user_agent": func() (bool, error) {
			return validator.Validate(target.EventData.UserAgent).MaxLength(1023).Regex("^[A-Za-z -]+$").GetResult()
		},
		"message": func() (bool, error) {
			return validator.Validate(target.EventData.Message).Required().MaxLength(8000).GetResult()
		},
		"datetime": func() (bool, error) {
			return validator.Validate(target.EventData.Datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		},
	}
	return mapValidationRules
}
