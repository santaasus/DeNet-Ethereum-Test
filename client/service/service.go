package client

import (
	pb "DeNet/pb/generated"
	generator "DeNet/utils"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

type IService interface {
	GetAccount()
	GetAccounts()
}

type Service struct {
	pb.AccountServiceClient
}

type streamType grpc.BidiStreamingClient[pb.AccountsRequest, pb.AccountsResponse]

func (s *Service) GetAccount() {
	req := &pb.AccountRequest{
		EthereumAddress: "0x1cE4870430739c947C716C1DF84cee7dD9A8B227",
		CryptoSignature: "7909ab6821ad46cf65516ccfd5183809bcd26b872d38ed3b930576ccab02026c1664b6de04102a2dca24d346f813ab432189c755fcc676eecd958b1d07205a9c00",
	}
	res, err := s.AccountServiceClient.GetAccount(context.Background(), req)
	if err != nil {
		log.Fatalf("GetAccount error: %v", err)
	}

	fmt.Printf("GetAccount response: Gastoken balance: %s, Wallet nonce: %s\n", res.GastokenBalance, res.WalletNonce)
}

func (s *Service) GetAccounts() {
	addressesChan := make(chan []string, 1)
	go func() {
		addresses, err := generator.Generate(100)
		if err != nil {
			log.Fatalf("%v", err)
		}

		addressesChan <- addresses
	}()

	addresses := <-addressesChan

	// Put yourself erc20 token's address
	token := "0xTokenAddress1"

	stream, err := s.AccountServiceClient.GetAccounts(context.Background())
	if err != nil {
		log.Fatalf("GetAccounts stream error: %v", err)
	}

	go streamToAccounts(stream, addresses, token)
	go streamFromAccounts(stream)
}

func streamToAccounts(stream streamType, addresses []string, tokenAddress string) {
	req := &pb.AccountsRequest{
		EthereumAddresses: addresses,
		Erc20TokenAddress: tokenAddress,
	}
	if err := stream.Send(req); err != nil {
		log.Fatalf("Failed to send: %v", err)
	}

	stream.CloseSend()
}

func streamFromAccounts(stream streamType) {
	startTime := time.Now()
	obtainedAddrsCounter := 0
	for {
		res, err := stream.Recv()
		obtainedAddrsCounter += 1

		if err != nil {
			continue
		}

		fmt.Printf("GetAccounts response: Address %s, Balance %s\n", res.EthereumAddresses, res.Erc20Balance)

		if obtainedAddrsCounter == len(res.EthereumAddresses)-1 {
			fmt.Printf("Elapsed time: %s\n", time.Since(startTime))

			break
		}
	}
}
