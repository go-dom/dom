package lottery

import (
	"context"

	"github.com/3JoB/ethclient"
)

func (c *Config) GetBlockHash() error {
	url := "https://eth.rpc.rivet.cloud/"
	if NodeUrl != "" {
		url = NodeUrl
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	defer client.Close()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	c.blockhash = header.Hash().Hex()
	return nil
}
