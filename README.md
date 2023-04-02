# go-lottery

A lottery algorithm library based on the Ethereum mainnet in Golang.

Usage:
```
go get gopkg.in/go-dom/lottery.v1
```

Example usage:
```
package main

import (
	"gopkg.in/go-dom/lottery.v1"
	"fmt"
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
	d, err := lottery.New(lottery.Config{
		ETHUrl: "https://326d90b517a047aa9be8cc4770d2a871.eth.rpc.rivet.cloud/",
		Lotteryid: "dgub8v7bvc7",
		UserNum: len(user),
		PrizeNum: 2,
		UserID: user,
	})
	if err!= nil {
        panic(err)
    }
	fmt.Println(d)
}
```