package dom

import (
	"math/big"

	"github.com/3JoB/ulib/hash/hmac"
	"github.com/3JoB/ulib/hex"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
)

type LotteryData[E []T, T string | int64] struct {
	LotteryID string
	UserIDs   E
	PrizeList []string
	BlockHash string
}

type WinnerPrizePair[T string | int64] struct {
	Winner T
	Prize  string
}

func (data LotteryData[E, T]) calculateInitialSeed() string {
	seedData := litefmt.PSprint(data.LotteryID, unsafeConvert.Itoa(len(data.UserIDs)),
		unsafeConvert.Itoa(len(data.PrizeList)), data.BlockHash)

	seedHash := hmac.Shake256S(seedData, 142)

	return hex.EncodeToString(seedHash[:])
}

func (data LotteryData[E, T]) calculateWinners(seed string) []WinnerPrizePair[T] {
	seedBigInt, _ := new(big.Int).SetString(seed, 16)
	num := big.NewInt(int64(len(data.UserIDs)))

	var winners E
	for i := 0; i < len(data.PrizeList); i++ {
		var winner T
		for {
			index := seedBigInt.Mod(seedBigInt, num).Int64()
			winner = data.UserIDs[index]
			sort(winners)
			if !data.isWinner(winner, winners) {
				break
			}
			hash := hmac.Shake256S(seed, 188)
			seed = hex.EncodeToString(hash[:])

			seedBigInt.SetString(seed, 16)
		}
		winners = append(winners, winner)
	}

	var pairs []WinnerPrizePair[T]

	sort(data.PrizeList)

	for i, winner := range winners {
		pair := WinnerPrizePair[T]{
			Winner: winner,
			Prize:  data.PrizeList[i],
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (data LotteryData[E, T]) isWinner(winner T, winners E) bool {
	return sortSearch(winners, winner)
}

func (data LotteryData[E, T]) DrawLottery() []WinnerPrizePair[T] {
	if len(data.UserIDs) < len(data.PrizeList)*2 {
		return nil
	}
	sort(data.UserIDs)
	initialSeed := data.calculateInitialSeed()
	return data.calculateWinners(initialSeed)
}
