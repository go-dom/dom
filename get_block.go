package lottery

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockHash() (string, error) {
	url := "https://api.ethereum.io/v1/eth/blockNumber"
	//"https://beaconstate.ethstaker.cc/"
	if NodeUrl != "" {
		url = NodeUrl
	}
	client, err := ethclient.Dial(url)
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
