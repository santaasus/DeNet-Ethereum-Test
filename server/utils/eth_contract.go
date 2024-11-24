package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"

	contract "DeNet/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// Tests naming for simple visualization
	FirstTokenID  = "DNT"
	SecondTokenID = "DNT2"
	ThirdTokenID  = "DNT2"
)

func DeployContract(erc20TokenAddress common.Address, ethclient *ethclient.Client) (multiToken *contract.MultiToken, err error) {
	// Deploy might be for an random user
	privateKeyHex := generatePrivateKey()
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Error for map the publicKey to the publicKeyECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	// If it doesn't needs transfer some funds
	auth.NoSend = false
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	block, err := ethclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		return
	}

	auth.GasLimit = block.GasLimit()

	if erc20TokenAddress.Hex() != "" {
		multiToken, err = contract.NewMultiToken(erc20TokenAddress, ethclient)
	} else {
		erc20TokenAddress, _, multiToken, err = contract.DeployMultiToken(auth, ethclient)
	}

	if err != nil {
		return nil, err
	}

	generateMultiTokens(multiToken)

	fmt.Printf("Contract succesfully deployed by address: %s\n", erc20TokenAddress)

	return multiToken, nil
}

func generateMultiTokens(contractToken *contract.MultiToken) {
	tokens := map[string]string{
		"DeNetToken":  FirstTokenID,
		"DeNetToken2": SecondTokenID,
		"DeNetToken3": ThirdTokenID,
	}

	for k, v := range tokens {
		_, err := contractToken.CreateToken(&bind.TransactOpts{}, k, v, big.NewInt(1000))
		if err != nil {
			fmt.Printf("Error for CreateToken %s", k)
		}
	}
}

func generatePrivateKey() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Ошибка генерации приватного ключа: %v", err)
	}

	privateKeyHex := hex.EncodeToString(crypto.FromECDSA(privateKey))

	return privateKeyHex
}
