package lottery

import (
	"fmt"

	"github.com/3JoB/ulib/crypt/hmac"
)

func (c *Config) Seeds() {
	c.seed = hmac.SHA512(fmt.Sprintf("%v:%v:%v", c.Lotteryid, c.UserNum, c.PrizeNum), c.blockhash)
}

func (c *Config) ReSeed() {
	c.seed = hmac.SHA512(c.seed, c.blockhash)
}
