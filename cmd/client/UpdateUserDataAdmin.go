package main

import (
	"context"
	"fmt"
	sorm "github.com/alexdsayandexru/sorm/gen"
)

func UpdateUserDataAdmin(client sorm.UserDataManagementClient) {
	response, err := client.UpdateUserDataAdmin(context.Background(),
		&sorm.UpdateUserDataAdminRequest{
			CorrelationId:   "01548f2c-dabf-11ee-a506-0242ac120002",
			TelcoId:         99,
			UserType:        100,
			UserId:          "qweRTY-_1234567890",
			ServiceId:       100000000,
			Datetime:        "05-03-2024:23.23.23.003",
			Msisdns:         []*sorm.Msisdn{{Msisdn: "+79261234567"}},
			DatetimeReg:     "05-03-2024:23.23.23.003",
			Emails:          []*sorm.Email{{Email: "qwerty@mail.ru"}},
			DatetimeUpdated: "05-03-2024:23.23.23.003",
			ServiceUser:     1,
			Additional:      []*sorm.Add{{Title: "", Content: ""}},
			ContractDate:    "05-03-2024:23.23.23.003",
			Nick:            "qwertyйцукенг",
			BirthDate:       "05-03-2024",
			Family:          "qwertyйцукенг",
			Initial:         "qwertyйцукенг",
			Address:         "qwertx yйцуvd zfvdzvaefcd  gкенг",
			ImIds:           []*sorm.ImId{{ServiceId: "ssss", ServiceName: "sssxsxa"}},
		})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
}
