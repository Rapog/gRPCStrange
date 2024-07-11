package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	team00v1 "teamclient/api/protos/gen/go/gRPCServer"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("New client err: ", err.Error())
	}

	defer conn.Close()

	client := team00v1.NewEx00Client(conn)
	ctx := context.Background()
	in := emptypb.Empty{}
	req, err := client.Connect(ctx, &in)

	if err != nil {
		log.Println("Connect err: ", err.Error())
	}
	for {
		stream, _ := req.Recv()
		//if err != io.EOF {
		//	return
		//} else if err == nil {
		fmt.Println(stream.SessionId)
		fmt.Println(stream.Frequency)
		fmt.Println(stream.Time.AsTime().UTC())
		fmt.Println()
		time.Sleep(1 * time.Second)
		//}
	}
}
