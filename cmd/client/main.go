package main

import (
	"context"
	"fmt"
	"github.com/alexdsayandexru/sorm/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := sorm.NewUserDataManagementClient(conn)

	response, err := client.RegisterUser(context.Background(), &sorm.RegisterUserRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
