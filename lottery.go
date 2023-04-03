package lottery

import (
	"fmt"

	"github.com/3JoB/ethclient/rpc"
)

type Config struct {
	// Example: https://eth.rpc.rivet.cloud/apikey
	APIUrl string
	ClientOption *rpc.ClientOption

	Lotteryid         string
	UserNum, PrizeNum int
	UserID            []int64
	Debug             bool

	blockhash string
	seed string
}

var NodeUrl string

func New(conf *Config) ([]int64, error) {
	if conf.APIUrl != "" {
		NodeUrl = conf.APIUrl
	}
	userhashs := BuildHash64(conf.UserID, conf.Lotteryid)
	userlist := IDS(userhashs)
	if err := conf.GetBlockHash(); err != nil {
		return nil, err
	}

	conf.Seeds()
	winners := conf.GetUser()

	var winnersID []int
	for _, winner := range winners {
		for i, userID := range userlist {
			if userID == winner {
				winnersID = append(winnersID, i)
				break
			}
		}
	}

	the_winners := make([]int64, 0, conf.PrizeNum)
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
	if conf.Debug {
		fmt.Printf("UserHashs: %v\n", userhashs)
		fmt.Printf("UserList: %v\n", userlist)
		fmt.Printf("BlockHash: %v\n", conf.blockhash)
		fmt.Printf("Seed: %v\n", conf.seed)
		fmt.Printf("Winners: %v\n", winners)
		fmt.Printf("WinnersID: %v\n", winnersID)
		fmt.Printf("TheWinners: %v\n", the_winners)
	}

	return the_winners, nil
}
