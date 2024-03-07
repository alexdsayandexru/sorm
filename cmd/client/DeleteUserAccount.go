package main

import (
	"context"
	"fmt"
	sorm "github.com/alexdsayandexru/sorm/gen"
)

func DeleteUserAccount(client sorm.UserDataManagementClient) {
	response, err := client.DeleteUserAccount(context.Background(),
		&sorm.DeleteUserAccountRequest{
			CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
			TelcoId:       99,
			UserType:      100,
			UserId:        "qweRTY-_1234567890",
			ServiceId:     100000000,
			Ip:            "127.0.0.1",
			Port:          99999,
			UserAgent:     "qwerty QWERTY-",
			Datetime:      "05-03-2024:23.23.23.003",
			Message:       "1234567890QWERTYqwerty",
		})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
}
