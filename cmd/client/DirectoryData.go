package main

import (
	"context"
	"fmt"
	sorm "github.com/alexdsayandexru/sorm/gen"
)

func DirectoryData(client sorm.UserDataManagementClient) {
	response, err := client.DirectoryData(context.Background(),
		&sorm.DirectoryDataRequest{
			CorrelationId: "01548f2c-dabf-11ee-a506-0242ac120002",
			Datetime:      "05-03-2024:23.23.23.003",
			Services: []*sorm.Service{
				{ServiceId: 1, Decription: "=hello=вцвцац", BeginTime: "05-03-2024:23.23.23.003", EndTime: "05-03-2024:23.23.23.003"},
				{ServiceId: 1, Decription: "=hello=вцвцац", BeginTime: "05-03-2024:23.23.23.003"},
				{ServiceId: 1, Decription: "=hello=вцвцац", BeginTime: "05-03-2024:23.23.23.003", EndTime: "05-03-2024:23.23.23"},
			},
		})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response)
}
