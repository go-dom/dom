package dom

import (
	"time"

	"github.com/3JoB/ulib/hash/hmac"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/ulid"
	"github.com/google/uuid"
	"lukechampine.com/frand"
)

type DBroke[E []T, T string | int64] struct {
	data *LotteryData[E, T]
}

func NewString() *DBroke[[]string, string] {
	return &DBroke[[]string, string]{
		data: &LotteryData[[]string, string]{
			PrizeList: make([]string, 1),
			UserIDs:   make([]string, 2),
		},
	}
}

func NewInt64() *DBroke[[]int64, int64] {
	return &DBroke[[]int64, int64]{
		data: &LotteryData[[]int64, int64]{
			PrizeList: make([]string, 1),
			UserIDs:   make([]int64, 2),
		},
	}
}

func (b *DBroke[E, T]) NewLotteryID() {
	uid, err := ulid.New(ulid.Timestamp(time.Now()), frand.New())
	id := ""
	if err != nil {
		id = uuid.NewString()
	} else {
		id = uid.String()
	}
	b.SetLotteryID(id)
}

func (b *DBroke[E, T]) SetUsers(data E) *DBroke[E, T] {
	b.data.UserIDs = data
	return b
}

func (b *DBroke[E, T]) SetBlock(data string) *DBroke[E, T] {
	b.data.BlockHash = litefmt.PSprint("0x", hmac.Shake128S(data, 64))
	return b
}

func (b *DBroke[E, T]) SetLotteryID(data string) *DBroke[E, T] {
	b.data.LotteryID = litefmt.PSprint("2x", hmac.Shake128S(data, 64))
	return b
}

func (b *DBroke[E, T]) SetPrizes(data []string) *DBroke[E, T] {
	b.data.PrizeList = data
	return b
}

func (b *DBroke[E, T]) AddUser(data T) *DBroke[E, T] {
	b.data.UserIDs = append(b.data.UserIDs, data)
	return b
}

func (b *DBroke[E, T]) AddPrize(data string) *DBroke[E, T] {
	b.data.PrizeList = append(b.data.PrizeList, data)
	return b
}

func (b *DBroke[E, T]) ExportData() *LotteryData[E, T] {
	return b.data
}

func (b *DBroke[E, T]) Do() ([]WinnerPrizePair[T], bool) {
	pair := b.data.DrawLottery()
	if len(pair) < 1 {
		return nil, false
	}
	return pair, true
}
