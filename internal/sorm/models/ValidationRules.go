package models

import "github.com/alexdsayandexru/sorm/internal/validator"

func ValidateMsisdnLogin(msisdn_login string, email_login string) (bool, error) {
	return validator.Validate(msisdn_login).RequiredIf(len(email_login) == 0).MaxLength(16).Regex("^[0-9+]+$").GetResult()
}

func ValidateEmailLogin(email_login string, msisdn_login string) (bool, error) {
	return validator.Validate(email_login).RequiredIf(len(msisdn_login) == 0).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
}

func ValidateFactor(factor string) (bool, error) {
	return validator.Validate(factor).MaxLength(255).Regex("^[A-Za-z0-9-]+$").GetResult()
}

func ValidateIp(ip string) (bool, error) {
	return validator.Validate(ip).MaxLength(255).Regex("^[0-9.]+$").GetResult()
}

func ValidatePort(port int32) (bool, error) {
	return validator.Validate(port).Maximum(99999).GetResult()
}

func ValidateUserAgent(user_agent string) (bool, error) {
	return validator.Validate(user_agent).MaxLength(1023).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
}

func ValidateMessage(message string) (bool, error) {
	return validator.Validate(message).Required().MaxLength(8000).GetResult()
}

func ValidateCorrelationId(correlation_id string) (bool, error) {
	return validator.Validate(correlation_id).Required().Uiid().GetResult()
}

func ValidateTelcoId(telco_id int32) (bool, error) {
	return validator.Validate(telco_id).Required().Maximum(100).GetResult()
}

func ValidateUserType(user_type int32) (bool, error) {
	return validator.Validate(user_type).Required().Maximum(100).GetResult()
}

func ValidateUserId(user_id string) (bool, error) {
	return validator.Validate(user_id).Required().MaxLength(255).Regex("^[A-Za-z0-9_-]+$").GetResult()
}

func ValidateNick(nick string) (bool, error) {
	return validator.Validate(nick).MaxLength(150).GetResult()
}

func ValidateBirthDate(birth_date string) (bool, error) {
	return validator.Validate(birth_date).Length(10).Regex("^[0-9-]+$").GetResult()
}

func ValidateNameFamilyInitial(name string) (bool, error) {
	return validator.Validate(name).MaxLength(30).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
}

func ValidateAddress(address string) (bool, error) {
	return validator.Validate(address).MaxLength(255).GetResult()
}

func ValidateRequiredDatetime(datetime string) (bool, error) {
	return validator.Validate(datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
}

func ValidateServiceUser(service_user int32) (bool, error) {
	return validator.Validate(service_user).Required().Equal(1).GetResult()
}

func ValidateServiceId(service_id int32) (bool, error) {
	return validator.Validate(service_id).Required().Maximum(100000000).GetResult()
}

func ValidateAdditional(additionals []Additional) (bool, error) {
	if additionals == nil {
		return true, nil
	}

	for i, m := range additionals {
		ok, err := validator.Validate(m.Title).MaxLength(255).Regex("^[A-Za-zА-Яа-я -]+$").GetResult()
		if !ok {
			return false, NewError(i, "title", err)
		}
		ok, err = validator.Validate(m.Content).MaxLength(255).GetResult()
		if !ok {
			return false, NewError(i, "content", err)
		}
	}
	return true, nil
}

func ValidateMsisdns(msisdns []string, emails []string) (bool, error) {
	ok, err := validator.Validate(msisdns).RequiredIf(emails == nil || len(emails) == 0).GetResult()
	if ok {
		for i, m := range msisdns {
			ok, err := validator.Validate(m).MaxLength(16).Regex("^[0-9+]+$").GetResult()
			if !ok {
				return false, NewError(i, "msisdn", err)
			}
		}
	}
	return ok, err
}

func ValidateEmails(emails []string, msisdns []string) (bool, error) {
	ok, err := validator.Validate(emails).RequiredIf(msisdns == nil || len(msisdns) == 0).GetResult()
	if ok {
		for i, m := range emails {
			ok, err := validator.Validate(m).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult()
			if !ok {
				return false, NewError(i, "email", err)
			}
		}
	}
	return ok, err
}

func ValidateImId(im_ids []ImId) (bool, error) {
	if im_ids == nil {
		return true, nil
	}

	for i, m := range im_ids {
		ok, err := validator.Validate(m.ServiceId).MaxLength(255).GetResult()
		if !ok {
			return false, NewError(i, "service_id", err)
		}
		ok, err = validator.Validate(m.ServiceName).MaxLength(255).GetResult()
		if !ok {
			return false, NewError(i, "service_name", err)
		}
	}
	return true, nil
}

func ValidateOriServices(oriServices []Service) (bool, error) {
	for i, s := range oriServices {
		ok, err := validator.Validate(s.ServiceId).Required().Maximum(100000000).GetResult()
		if !ok {
			return false, NewError(i, "service_id", err)
		}
		ok, err = validator.Validate(s.Description).MaxLength(255).GetResult()
		if !ok {
			return false, NewError(i, "description", err)
		}
		ok, err = validator.Validate(s.BeginTime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult()
		if !ok {
			return false, NewError(i, "begin_time", err)
		}
		ok, err = validator.Validate(s.EndTime).Length(23).Regex("^[0-9 :.-]+$").GetResult()
		if !ok {
			return false, NewError(i, "end_time", err)
		}
	}
	return true, nil
}
