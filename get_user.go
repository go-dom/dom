package lottery

import (
	"math/big"
)

func (c *Config) GetUser() []int64 {
	bigSeed, _ := new(big.Int).SetString(c.seed, 16)
	var winner []int64
	for i := 0; i < c.PrizeNum; i++ {
		winnerID := bigSeed.Mod(bigSeed, big.NewInt(int64(c.UserNum))).Int64()
		if winnerID != 0 {
			winner = append(winner, winnerID)
			c.ReSeed()
			bigSeed, _ = new(big.Int).SetString(c.seed, 16)
		}
	}
	return winner
}
