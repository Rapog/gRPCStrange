package ex00

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	team00v1 "server/api/protos/gen/go/gRPCServer"
	"server/internal/services/srvStream"
)

type serverAPI struct {
	team00v1.UnimplementedEx00Server
	streamService srvStream.StreamProvider
}

func Register(gRPC *grpc.Server) {
	team00v1.RegisterEx00Server(gRPC, &serverAPI{})
}

func (s *serverAPI) Connect(ctx context.Context, emp *emptypb.Empty) (*team00v1.ConnectResponse, error) {
	//fmt.Println("model.SessionId")
	s.streamService = srvStream.NewStrmPrvder()
	model, _ := s.streamService.Stream(ctx)
	//fmt.Printf(model.SessionId)
	return &team00v1.ConnectResponse{
		SessionId: model.SessionId,
		Frequency: model.Frequency,
		Time:      timestamppb.New(model.Time),
	}, nil
	//return &team00v1.ConnectResponse{
	//	SessionId: "123",
	//	Frequency: 1.123,
	//	Time:      timestamppb.New(time.Now().UTC()),
	//}, nil
}
