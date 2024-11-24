package adapter

import (
	pb "DeNet/pb/generated"
	"DeNet/server/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type Adapter struct {
	*grpc.Server
	*ethclient.Client
}

func (a *Adapter) AccountAdapter() {
	service := &service.Service{Client: a.Client}
	pb.RegisterAccountServiceServer(a.Server, service)
}
