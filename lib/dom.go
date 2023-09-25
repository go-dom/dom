package dom

type DBroke[E []T, T string | int64] struct {
	data *LotteryData[E, T]
}

func NewString() *DBroke[[]string, string] {
	return &DBroke[[]string, string]{
		data: &LotteryData[[]string, string]{
			PrizeList: make([]string, 1),
			UserIDs: make([]string, 2),
		},
	}
}

func NewInt64() *DBroke[[]int64, int64] {
	return &DBroke[[]int64, int64]{
		data: &LotteryData[[]int64, int64]{
			PrizeList: make([]string, 1),
			UserIDs: make([]int64, 2),
		},
	}
}

func (b *DBroke[E, T]) SetUsers(data E) *DBroke[E, T] {
	b.data.UserIDs = data
	return b
}

func (b *DBroke[E, T]) SetBlock(data string) *DBroke[E, T] {
	b.data.BlockHash = data
	return b
}

func (b *DBroke[E, T]) SetLotteryID(data string) *DBroke[E, T] {
	b.data.LotteryID = data
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

func (b *DBroke[E, T]) Do() ([]WinnerPrizePair[T], bool) {
	pair := b.data.DrawLottery()
	if len(pair) < 1 {
		return nil, false
	}
	return pair, true
}