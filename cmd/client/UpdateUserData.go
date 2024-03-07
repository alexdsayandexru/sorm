package main

import (
	"context"
	"fmt"
	sorm "github.com/alexdsayandexru/sorm/gen"
)

func UpdateUserData(client sorm.UserDataManagementClient) {
	response, err := client.UpdateUserData(context.Background(),
		&sorm.UpdateUserDataRequest{
			CorrelationId:   "01548f2c-dabf-11ee-a506-0242ac120002",
			TelcoId:         99,
			UserType:        100,
			UserId:          "qweRTY-_1234567890",
			ServiceId:       100000000,
			Ip:              "127.0.0.1",
			Port:            99999,
			UserAgent:       "рлопрлорлор qwerty QWERTY-",
			Datetime:        "05-03-2024:23.23.23.003",
			Msisdns:         []*sorm.Msisdn{{Msisdn: "+79261234567"}},
			Emails:          []*sorm.Email{{Email: "qwerty@mail.ru"}},
			DatetimeReg:     "05-03-2024:23.23.23.003",
			DatetimeUpdated: "05-03-2024:23.23.23.003",
			ServiceUser:     1,
			Additional:      []*sorm.Add{{Title: "порпо", Content: "hk"}},
			ContractDate:    "05-03-2024:23.23.23.003",
			Nick:            "qwertyuiop1234567890паавиа",
			BirthDate:       "05-03-2024",
			Name:            "qwertyuiopЙЦУКЕН",
			Family:          "qwertyuiopйцукен",
			Initial:         "qwertyuiopкуцйы",
			Address:         "qwertyuiop1234567890",
			ImIds:           []*sorm.ImId{{ServiceId: "hello", ServiceName: "world"}},
			Message:         "qwertyuiop1234567890",
		})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
}
