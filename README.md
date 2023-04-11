# go-dom

A lottery algorithm library based on the Ethereum mainnet in Golang.

1.5+

Usage:
```sh
go get gopkg.in/dom.v1
```

1.5-
```sh
go get gopkg.in/go-dom/lottery.v1
```

Example usage:
```golang
package main

import (
	"fmt"

	"gopkg.in/dom.v1"
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
	client, err := dom.NewClient().SetUrl("https://apikey.eth.rpc.rivet.cloud/").SetDebug().Dial()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	session, err := client.NewSession(&lottery.Session{
		UserNum:  len(user),
		PrizeNum: 2,
		UserID:   user,
	})
	if err != nil {
		panic(err)
	}

	data, err := session.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```

## Known issues

**These issues are only sporadic and require developers to verify themselves. For example, retry when a problem occurs.**

- Sometimes, some UserID will win continuously.
- Sometimes there will be fewer winners than the set number of prizes.


# License
This software is distributed under GNU Affero General Public License v3.0 license.