package ex00

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	team00v1 "server/api/protos/gen/go/gRPCServer"
	"server/internal/cache"
	"server/internal/domain/models"
	"server/internal/services/srvStream"
	"time"
)

type serverAPI struct {
	team00v1.UnimplementedEx00Server
	streamService srvStream.StreamProvider
	cache.Cache[*models.MeanStd]
}

func Register(gRPC *grpc.Server, cache cache.Cache[*models.MeanStd]) {
	team00v1.RegisterEx00Server(gRPC, &serverAPI{Cache: cache})
}

func (s *serverAPI) Connect(emp *emptypb.Empty, cnct team00v1.Ex00_ConnectServer) error {
	cche := cache.New[*models.MeanStd](10)
	s.streamService = srvStream.NewStrmPrvder(cche)
	ctx := context.Background()

	model := &models.SrvStream{}

	connectionUuid, _ := uuid.NewUUID()
	connectUuid := connectionUuid.String()

	for {
		model, _ = s.streamService.Stream(ctx, connectUuid)
		err := cnct.Send(&team00v1.ConnectResponse{
			SessionId: model.SessionId,
			Frequency: model.Frequency,
			Time:      timestamppb.New(model.Time),
		})
		//fmt.Println(model)
		if err != nil {
			cnct.SendMsg("Error: SendConnectionResponse: " + err.Error())
			break
		}
		time.Sleep(1 * time.Second)
	}

	//fmt.Printf(model.SessionId)
	//return &team00v1.ConnectResponse{
	//	SessionId: model.SessionId,
	//	Frequency: model.Frequency,
	//	Time:      timestamppb.New(model.Time),
	//}, nil
	//return &team00v1.ConnectResponse{
	//	SessionId: "123",
	//	Frequency: 1.123,
	//	Time:      timestamppb.New(time.Now().UTC()),
	//}, nil
	return nil
}
