package dom

import (
	"encoding/hex"
	"math/big"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/sha3"
)

type LotteryData struct {
	LotteryID string
	UserIDs   []string
	PrizeList []string
	BlockHash string
}
type WinnerPrizePair struct {
	Winner string
	Prize  string
}

func calculateInitialSeed(data LotteryData) string {
	seedData := data.LotteryID + unsafeConvert.Itoa(len(data.UserIDs)) +
		unsafeConvert.Itoa(len(data.PrizeList)) + data.BlockHash

	sha := sha3.New512()
	sha.Write([]byte(seedData))
	seedHash := sha.Sum(nil)

	return hex.EncodeToString(seedHash[:])
}

func calculateWinners(seed string, data LotteryData) []WinnerPrizePair {
	seedBigInt, _ := new(big.Int).SetString(seed, 16)
	num := big.NewInt(int64(len(data.UserIDs)))

	var winners []string
	for i := 0; i < len(data.PrizeList); i++ {
		index := seedBigInt.Mod(seedBigInt, num).Int64()
		winner := data.UserIDs[index]
		winners = append(winners, winner)
		sha := sha3.New512()
		sha.Write([]byte(seed))
		hash := sha.Sum(nil)
		seed = hex.EncodeToString(hash[:])

		seedBigInt.SetString(seed, 16)
	}

	var pairs []WinnerPrizePair

	for i, winner := range winners {
		pair := WinnerPrizePair{
			Winner: winner,
			Prize:  data.PrizeList[i],
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

func DrawLottery(data LotteryData) []WinnerPrizePair {
	sortUserIDs(data.UserIDs)
	initialSeed := calculateInitialSeed(data)
	return calculateWinners(initialSeed, data)
}
