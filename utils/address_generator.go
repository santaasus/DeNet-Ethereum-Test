package utils

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

func Generate(count uint) ([]string, error) {
	var addresses []string
	for range 100 {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("Error for cast to ecdsa.PublicKey type")
		}

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		addresses = append(addresses, address)
	}

	return addresses, nil
}
