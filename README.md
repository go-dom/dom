# go-lottery

A lottery algorithm library based on the Ethereum mainnet in Golang.

Usage:
```sh
go get gopkg.in/go-dom/lottery.v1
```

Example usage:
```golang
package main

import (
	"fmt"

	"gopkg.in/go-dom/lottery.v1"
)

func main() {
	user := []int64{
		124875175,
		12848475,
		15768612,
		432867286,
		3268742,
		262274327,
		27923727382,
		23672472472,
		72472472,
	}
	client, err := lottery.NewClient().SetUrl("https://apikey.eth.rpc.rivet.cloud/").SetDebug().Dial()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	data, err := client.NewStream(&lottery.Data{
		UserNum:  len(user),
		PrizeNum: 2,
		UserID:   user,
	}).Do()

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```