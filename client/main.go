package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"math"
	"sync"
	team00v1 "teamclient/api/protos/gen/go/gRPCServer"
	"teamclient/internal/meanStdFunc"
	"teamclient/internal/repository/postgresGorm"
	"time"
)

func main() {

	koefK := flag.Float64("k", 0, "koef of frequency's STD")
	flag.Parse()

	if *koefK == 0 {
		log.Fatalf("Flag -k didnt setted")
	}

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
	storage, err := postgresGorm.NewDB()
	if err != nil {
		log.Fatalln("storage.NewDB: ", err.Error())
	}

	var count int
	var stdSlice = make([]float64, 0, 50)
	var stdev, mean, expdVal float64
	var stdevcalcd bool

	var freqPoll = sync.Pool{
		New: func() interface{} {
			return new(*team00v1.ConnectResponse)
		},
	}
	for {
		stream := freqPoll.Get().(**team00v1.ConnectResponse)
		*stream, _ = req.Recv()
		if !stdevcalcd && len(stdSlice) < 50 {
			stdSlice = append(stdSlice, (*stream).Frequency)
		} else {
			mean = meanStdFunc.MeanFunc(stdSlice)
			stdev, stdevcalcd = meanStdFunc.STDDevFunc(stdSlice, mean)
			expdVal = stdev * *koefK
			//fmt.Println("Mean: ", mean)
			//fmt.Println("STD: ", stdev)
			//fmt.Println("ExpdVal", expdVal)
			break
		}
		freqPoll.Put(stream)
	}

	for {
		stream, _ := req.Recv()
		//fmt.Println(stream.Frequency)
		count++
		if math.Abs(stream.Frequency-mean) > math.Abs(expdVal) {
			fmt.Println("Anomaly detected: Values #", count, "Value: ", stream.Frequency)
			err := storage.AddAnomaly(*stream)
			if err != nil {
				log.Println("Adding anomaly err: ", err)
			}
			time.Sleep(1 * time.Second)
		}
		//if err != io.EOF {
		//	return
		//} else if err == nil {
		//fmt.Println(stream.SessionId)
		//fmt.Println(stream.Frequency)
		//fmt.Println(stream.Time.AsTime().UTC())
		//fmt.Println()
		//time.Sleep(1 * time.Second)
		//}
		if count%10 == 0 {
			fmt.Println("Numbers received: ", count)
		}
	}
}

//docker run --name=anom_bd -e POSTGRES_PASSWORD='1234' -p 5432:5432 -d --rm postgres
//migrate -path ./schema -database 'postgres://postgres:1234@localhost:5432/postgres?sslmode=disable' up
