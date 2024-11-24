package server

import (
	"DeNet/server/adapter"
	"log"
	"net"

	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

func handleHosts() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ethClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}

	server := grpc.NewServer()

	adapter := adapter.Adapter{Client: ethClient}
	adapter.AccountAdapter()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	handleHosts()
}
