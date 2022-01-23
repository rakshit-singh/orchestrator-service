package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	m "github.com/rakshit-singh/orchestrator-service/DataMock"
	grpc "google.golang.org/grpc"
)

type database struct {
	*m.UnimplementedDataServiceServer
}

func (*database) GetMockUserData(ctx context.Context, name *m.UserNameDS) (*m.UserDS, error) {
	nameStr := name.GetName() //variable containing the name given by the client

	if len([]rune(nameStr)) < 6 {
		return nil, errors.New("the name is too short")
	} else {
		usr := &m.UserDS{Name: nameStr,
			Class: strconv.Itoa(len(nameStr)),
			Roll:  10 * int32((len(nameStr)))}

		return usr, nil
	}

}

func main() {

	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opt []grpc.ServerOption
	grpcServer := grpc.NewServer(opt...)

	var server m.DataServiceServer = &database{}

	m.RegisterDataServiceServer(grpcServer, server)

	grpcServer.Serve(lis)
}
