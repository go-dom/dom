package lottery

import (
	"fmt"

	"github.com/3JoB/ulib/crypt/hmac"
)

func Seeds(lotteryid, blockhash string, userNum, prizeNum int) string {
	return hmac.SHA512(fmt.Sprintf("%v:%v:%v", lotteryid, userNum, prizeNum), blockhash)
}

func ReSeed(seed, blockhash string) string {
	return hmac.SHA512(seed, blockhash)
}