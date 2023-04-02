package lottery

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockHash() (string, error) {
	client, err := ethclient.Dial("https://beaconstate.ethstaker.cc/")
	if err != nil {
		return "", err
	}
	defer client.Close()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", err
	}
	return header.Hash().Hex(), nil
}
