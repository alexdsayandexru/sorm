package main

import (
	"context"
	"fmt"
	sorm "github.com/alexdsayandexru/sorm/gen"
)

func RegisterUser(client sorm.UserDataManagementClient) {
	response, err := client.RegisterUser(context.Background(),
		&sorm.RegisterUserRequest{
			CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
			TelcoId:       99,
			UserType:      100,
			UserId:        "qweRTY-_1234567890",
			Msisdns:       []*sorm.Msisdn{{Msisdn: "+79261234567"}},
			Emails:        []*sorm.Email{{Email: "qwerty@mail.ru"}},
			DatetimeReg:   "05-03-2024:23.23.23.003",
			ServiceUser:   1,
			ContractDate:  "05-03-2024:23.23.23.003",
			ServiceId:     100000000,
			MsisdnLogin:   "+79261234567",
			EmailLogin:    "qwerty@mail.ru",
			Ip:            "127.0.0.1",
			Port:          99999,
			UserAgent:     "qwerty QWERTY-",
			Message:       "qwertyuiop1234567890",
			Datetime:      "05-03-2024:23.23.23.003",
			Additional:    []*sorm.Add{{Title: "", Content: ""}},
		})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
}
