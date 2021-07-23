package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vitorath/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connection to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
	AddUserVerbose(client)
	AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "v1",
			Name:  "Vitor",
			Email: "vitor@vitor.com",
		},
		&pb.User{
			Id:    "v2",
			Name:  "Vitor2",
			Email: "vitor2@vitor.com",
		},
		&pb.User{
			Id:    "v3",
			Name:  "Vitor3",
			Email: "vitor3@vitor.com",
		},
		&pb.User{
			Id:    "v4",
			Name:  "Vitor4",
			Email: "vitor4@vitor.com",
		},
		&pb.User{
			Id:    "v5",
			Name:  "Vitor5",
			Email: "vitor5@vitor.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Errror receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "v1",
			Name:  "Vitor",
			Email: "vitor@vitor.com",
		},
		&pb.User{
			Id:    "v2",
			Name:  "Vitor2",
			Email: "vitor2@vitor.com",
		},
		&pb.User{
			Id:    "v3",
			Name:  "Vitor3",
			Email: "vitor3@vitor.com",
		},
		&pb.User{
			Id:    "v4",
			Name:  "Vitor4",
			Email: "vitor4@vitor.com",
		},
		&pb.User{
			Id:    "v5",
			Name:  "Vitor5",
			Email: "vitor5@vitor.com",
		},
	}

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}

			fmt.Printf("Receiving user %v with states %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
