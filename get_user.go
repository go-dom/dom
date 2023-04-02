package lottery

import (
	"math/big"
)

func GetUser(seed, blockhash string, userNum, prizeNum int) []int64 {
	bigSeed, _ := new(big.Int).SetString(seed, 16)
	winner := make([]int64, 0, prizeNum)
	for i := 0; i < prizeNum; i++ {
		winnerID := bigSeed.Mod(bigSeed, big.NewInt(int64(userNum))).Int64()
		if winnerID != 0 {
			winner = append(winner, winnerID)
			seed = ReSeed(seed, blockhash)
			bigSeed, _ = new(big.Int).SetString(seed, 16)
		}
	}
	return winner
}