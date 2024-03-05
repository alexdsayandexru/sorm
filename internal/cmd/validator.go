package main

import (
	"fmt"
	"github.com/alexdsayandexru/sorm/internal/validator"
)

func ValidRegisterUser(
	correlation_id interface{}, telco_id interface{}, user_type interface{}, user_id interface{},
	msisdns interface{}, emails interface{}, datetime_reg interface{}, service_user interface{},
	title interface{}, content interface{}, contract_date interface{}, service_id interface{},
	msisdn_login interface{}, email_login interface{}, ip interface{}, port interface{},
	user_agent interface{}, message interface{}, datetime interface{}) {
	fmt.Println(validator.Validate(correlation_id).Required().Uiid().GetResult())
	fmt.Println(validator.Validate(telco_id).Required().Maximum(100).GetResult())
	fmt.Println(validator.Validate(user_type).Required().Maximum(100).GetResult())
	fmt.Println(validator.Validate(user_id).Required().MaxLength(255).Regex("^[A-Za-z0-9_-]+$").GetResult())
	fmt.Println(validator.Validate(msisdns).RequiredIf(emails == nil).MaxLength(16).Regex("^[0-9+]+$").GetResult())
	fmt.Println(validator.Validate(emails).RequiredIf(msisdns == nil).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult())
	fmt.Println(validator.Validate(datetime_reg).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult())
	fmt.Println(validator.Validate(service_user).Required().Equal(1).GetResult())
	fmt.Println(validator.Validate(title).MaxLength(255).Regex("^[A-Za-z_-]+$").GetResult())
	fmt.Println(validator.Validate(content).MaxLength(255).GetResult())
	fmt.Println(validator.Validate(contract_date).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult())
	fmt.Println(validator.Validate(service_id).Required().Maximum(100000000).GetResult())
	fmt.Println(validator.Validate(msisdn_login).RequiredIf(email_login == nil).MaxLength(16).Regex("^[0-9+]+$").GetResult())
	fmt.Println(validator.Validate(email_login).RequiredIf(msisdn_login == nil).MaxLength(100).Regex("^[A-Za-z0-9!#$%&‘*+/=?^_`{|}~@.-]+$").GetResult())
	fmt.Println(validator.Validate(ip).MaxLength(255).Regex("^[0-9.]+$").GetResult())
	fmt.Println(validator.Validate(port).Maximum(99999).GetResult())
	fmt.Println(validator.Validate(user_agent).MaxLength(1023).Regex("^[A-Za-z -]+$").GetResult())
	fmt.Println(validator.Validate(message).Required().MaxLength(8000).GetResult())
	fmt.Println(validator.Validate(datetime).Required().Length(23).Regex("^[0-9 :.-]+$").GetResult())
}

func main() {
	ValidRegisterUser("01548f2c-dabf-11ee-a506-0242ac120002", 99, 100, "qweRTY-_1234567890",
		"+79261234567", "qwerty@mail.ru", "05-03-2024:23.23.23.003", 1,
		"qwerty-QWERTY_", nil, "05-03-2024:23.23.23.003", 100000000,
		"+79261234567", "qwerty@mail.ru", "127.0.0.1", 99999,
		"qwerty QWERTY-", "qwertyuiop1234567890", "05-03-2024:23.23.23.003")

	/*b, err := json.Marshal(ru)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))*/
}
