package dom

import (
	"crypto/sha512"
	"encoding/hex"
	"math/big"
	"strconv"

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
	seedData := data.LotteryID + strconv.Itoa(len(data.UserIDs)) +
		strconv.Itoa(len(data.PrizeList)) + data.BlockHash

	seedHash := sha512.Sum512([]byte(seedData))

	return hex.EncodeToString(seedHash[:])
}

func calculateWinners(seed string, data LotteryData) []WinnerPrizePair {
	seedBigInt := new(big.Int)
	seedBigInt.SetString(seed, 16)

	var winners []string
	for i := 0; i < len(data.PrizeList); i++ {
		index := seedBigInt.Mod(seedBigInt, big.NewInt(int64(len(data.UserIDs)))).Int64()
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
