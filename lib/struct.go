package dom

import (
	"math/big"

	"github.com/3JoB/ulib/hex"
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
	sha.Write(unsafeConvert.ByteSlice(seedData))
	seedHash := sha.Sum(nil)

	return hex.EncodeToString(seedHash[:])
}

func calculateWinners(seed string, data LotteryData) []WinnerPrizePair {
	seedBigInt, _ := new(big.Int).SetString(seed, 16)
	num := big.NewInt(int64(len(data.UserIDs)))

	var winners []string
	for i := 0; i < len(data.PrizeList); i++ {
		winner := ""
		for {
			index := seedBigInt.Mod(seedBigInt, num).Int64()
			winner = data.UserIDs[index]
			if !isWinner(winner, winners) {
				break
			}
			sha := sha3.New512()
			sha.Write(unsafeConvert.ByteSlice(seed))
			hash := sha.Sum(nil)
			seed = hex.EncodeToString(hash[:])

			seedBigInt.SetString(seed, 16)
		}
		winners = append(winners, winner)
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

func isWinner(winner string, winners []string) bool {
	for _, w := range winners {
		if w == winner {
			return true
		}
	}
	return false
}

func DrawLottery(data LotteryData) []WinnerPrizePair {
	if len(data.UserIDs) < len(data.PrizeList)*2 {
		return nil
	}
	sortUserIDs(data.UserIDs)
	initialSeed := calculateInitialSeed(data)
	return calculateWinners(initialSeed, data)
}
