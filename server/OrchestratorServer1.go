package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	m "github.com/rakshit-singh/orchestrator-service"
	datamock "github.com/rakshit-singh/orchestrator-service/DataMock"
	"google.golang.org/grpc"
)

//contains the port number on which the current orchestrator-service is to be hosted on
var port string

type OrchestratorServer1 struct {
	*m.UnimplementedProcessRequestServer
}

// fetches the details of the user with name = UserName.getName()
func (o *OrchestratorServer1) GetUserByName(ctx context.Context, name *m.UserName) (*m.User, error) {

	// If the service is hosted on port 9000 call service on port 90001 & return the values recieved
	//If service is hosted on port 9001 call the datamock service on port 1000 & return values recieved
	switch port {
	case "9000":

		conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Failed to create a channel for communication %v", err.Error())
		}

		defer conn.Close()

		client := m.NewProcessRequestClient(conn)

		user, err1 := client.GetUserByName(ctx, name)

		return user, err1

	case "9001":
		return o.GetUser(ctx, name)
	default:
		return nil, errors.New("there is some problem with the ports on which servers are deployed")
	}
}

// This function calls the DataMockService to get the user data.
//Executed by the orchestrator service hosted on port 9001
func (*OrchestratorServer1) GetUser(ctx context.Context, name *m.UserName) (*m.User, error) {

	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to create a channel for communication %v", err.Error())
	}

	defer conn.Close()

	client := datamock.NewDataServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(90)*time.Second)

	defer cancel()

	//Converting userName type defined in OrchestratorService.proto to UserName type defined in DataMock.proto
	datamockName := datamock.UserNameDS{
		Name: name.GetName(),
	}
	user, err1 := client.GetMockUserData(ctx, &datamockName)

	return &m.User{Name: user.GetName(), Class: user.GetClass(), Roll: user.GetRoll()}, err1

}
func main() {

	args := os.Args

	port = args[1]

	lis, err := net.Listen("tcp", fmt.Sprint("localhost:", port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	var server m.ProcessRequestServer = &OrchestratorServer1{}

	m.RegisterProcessRequestServer(grpcServer, server)

	grpcServer.Serve(lis)
}
