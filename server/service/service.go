package service

import (
	"DeNet/contract"
	pb "DeNet/pb/generated"
	"DeNet/server/utils"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IService interface {
	GetAccount(ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error)
	GetAccounts(stream pb.AccountService_GetAccountsServer) error
}

type Service struct {
	pb.UnimplementedAccountServiceServer
	*ethclient.Client
}

func (s *Service) GetAccount(ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error) {
	address := req.EthereumAddress
	signature := req.CryptoSignature

	if !IsValidSignature(address, signature) {
		return nil, fmt.Errorf("Invalid signature")
	}

	res, err := s.getAccountBalance(address)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GetAccounts(stream pb.AccountService_GetAccountsServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		contract, err := utils.DeployContract(common.HexToAddress(req.Erc20TokenAddress), s.Client)
		if err != nil {
			log.Fatalf("Failed to DeployContract: %v", err)
		}

		for _, address := range req.EthereumAddresses {
			res, err := s.getAccountsBalance(contract, address)
			if err != nil {
				continue
			}

			if err := stream.SendMsg(res); err != nil {
				return err
			}
		}
	}
}

func (s *Service) getAccountBalance(address string) (*pb.AccountResponse, error) {
	cAddress := common.HexToAddress(address)
	balance, err := s.Client.BalanceAt(context.Background(), cAddress, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve balance: %v", err))
	}
	log.Printf("Token balance: %s", balance.String())

	nonce, err := s.Client.PendingNonceAt(context.Background(), cAddress)
	if err != nil {
		log.Fatalf("Error for getting PendingNonceAt: %v", err)
	}

	return &pb.AccountResponse{
		GastokenBalance: balance.String(),
		WalletNonce:     fmt.Sprintf("%d", nonce),
	}, nil
}

func (s *Service) getAccountsBalance(contract *contract.MultiToken, address string) (*pb.AccountsResponse, error) {
	balance, err := contract.BalanceOf(&bind.CallOpts{}, utils.FirstTokenID, common.HexToAddress(address))
	if err != nil {
		log.Fatalf("Failed to retrieve balance: %v", err)
	}
	log.Printf("Token balance: %s", balance.String())

	return &pb.AccountsResponse{
		EthereumAddresses: address,
		Erc20Balance:      balance.String(),
	}, nil
}

func IsValidSignature(address string, signature string) bool {
	// needs move to keystore
	message := "Hello, Ethereum!"
	hash := crypto.Keccak256Hash([]byte(message))
	pubKey, err := crypto.SigToPub(hash.Bytes(), common.Hex2Bytes(signature))
	if err != nil {
		return false
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	return recoveredAddress.Hex() == address
}
