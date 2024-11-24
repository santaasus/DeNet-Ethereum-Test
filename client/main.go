package main

import (
	"log"

	"DeNet/client/service"
	pb "DeNet/pb/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	pbClient := pb.NewAccountServiceClient(conn)
	service := &client.Service{
		AccountServiceClient: pbClient,
	}

	// service.GetAccount()
	service.GetAccounts()
}

func startEth() {

}
