package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	m "github.com/rakshit-singh/orchestrator-service"
	"google.golang.org/grpc"
)

var name string

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to create a channel for communication %v", err.Error())
	}

	defer conn.Close()

	client := m.NewProcessRequestClient(conn)

	name = os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(90)*time.Second)

	defer cancel()

	user, err1 := client.GetUserByName(ctx, createUserName(name))

	if err1 != nil {
		log.Fatalf("%v", err1.Error())
	} else {

		fmt.Println(user)
	}

}

//Wrapper function to create UserNAme object by taking a string as input argument
func createUserName(name string) *m.UserName {
	return &m.UserName{Name: name}
}
