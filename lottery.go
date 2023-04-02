package lottery

import "fmt"

type Config struct {
	// Example: https://eth.rpc.rivet.cloud/apikey
	ETHUrl string
	Lotteryid         string
	UserNum, PrizeNum int
	UserID            []int64
}

var NodeUrl string

func New(conf Config) ([]int64, error) {
	if conf.ETHUrl != "" {
		NodeUrl = conf.ETHUrl
	}
	userhashs := BuildHash64(conf.UserID, conf.Lotteryid)
	fmt.Printf("UserHashs: %v\n", userhashs)
	userlist := IDS(userhashs)
	fmt.Printf("UserList: %v\n", userlist)
	blockhash, err := GetBlockHash()
	if err != nil {
		return nil, err
	}
	fmt.Printf("BlockHash: %v\n", blockhash)
	seed := Seeds(conf.Lotteryid, blockhash, conf.UserNum, conf.PrizeNum)
	fmt.Printf("Seed: %v\n", seed)
	winners := GetUser(seed, blockhash, conf.UserNum, conf.PrizeNum)
	fmt.Printf("Winners: %v\n", winners)
	var winnersID []int
	for _, winner := range winners {
		for i, userID := range userlist {
			if userID == winner {
				winnersID = append(winnersID, i)
				break
			}
		}
	}
	fmt.Printf("WinnersID: %v\n", winnersID)

	the_winners := make([]int64, 0 , conf.PrizeNum)
	for _, winnerID := range winnersID {
		if Hash64(conf.UserID[winnerID], conf.Lotteryid) == userhashs[winnerID] {
			the_winners = append(the_winners, conf.UserID[winnerID])
		} else {
			for _, t := range conf.UserID {
				if Hash64(t, conf.Lotteryid) == userhashs[winnerID] {
					the_winners = append(the_winners, conf.UserID[winnerID])
					break
				}
			}
		}
	}
	fmt.Printf("TheWinners: %v\n", the_winners)
	return the_winners, nil
}
