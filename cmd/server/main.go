package main

import (
	"fmt"
	"github.com/alexdsayandexru/sorm/internal/grpc"
)

func main() {
	err := grpc.NewServer(9090).Run()
	if err != nil {
		fmt.Println(err)
	}
}
