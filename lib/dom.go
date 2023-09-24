package dom

type DBroke[E []T, T string | int64] struct {
	data *LotteryData[E, T]
}

func NewString() *DBroke[[]string, string] {
	return &DBroke[[]string, string]{
		data: &LotteryData[[]string, string]{},
	}
}

func NewInt64() *DBroke[[]int64, int64] {
	return &DBroke[[]int64, int64]{
		data: &LotteryData[[]int64, int64]{},
	}
}
