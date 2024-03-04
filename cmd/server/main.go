package main

import (
	"context"
	"fmt"
	"github.com/alexdsayandexru/sorm/gen"
	"google.golang.org/grpc"
	"net"
)

type DataCollectionServiceImpl struct {
	sorm.UnimplementedDataCollectionServiceServer
}

func (DataCollectionServiceImpl) Send(ctx context.Context, r *sorm.SendRequest) (*sorm.SendResponse, error) {
	fmt.Println("Send to Kafka:", r.Data)

	response := &sorm.SendResponse{Reserved: r.Data}
	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	server := grpc.NewServer()
	var dataCollectionServiceImpl DataCollectionServiceImpl
	sorm.RegisterDataCollectionServiceServer(server, dataCollectionServiceImpl)
	err = server.Serve(listen)
	if err != nil {
		fmt.Println(err)
		return
	}
}
