package lottery

type Config struct {
	Lotteryid string
	UserNum, PrizeNum int
	UserID []int64
}

func New(conf Config) ([]int64, error){
	userhashs := BuildHash64(conf.UserID, conf.Lotteryid)
	blockhash, err := GetBlockHash()
	if err != nil {
		return nil, err
	}
	seed := Seeds(conf.Lotteryid, blockhash, conf.UserNum, conf.PrizeNum)
	usid := GetUser(seed, blockhash, conf.UserNum, conf.PrizeNum)
	for i := range usid {
		usid[i] = userhashs[usid[i]]
	}
}