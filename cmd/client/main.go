package main

import (
	"fmt"
	"github.com/alexdsayandexru/sorm/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := sorm.NewUserDataManagementClient(conn)
	//RegisterUser(client)
	//LoginUser(client)
	//LogoutUser(client)
	//DeleteUserAccount(client)
	//UpdateUserData(client)
	//DeleteUserAccountAdmin(client)
	//UpdateUserDataAdmin(client)
	//DeleteAccount(client)
	//UserAccountRecovery(client)
	DirectoryData(client)
}
