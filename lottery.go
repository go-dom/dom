package lottery

import (
	"context"
	"fmt"
	"math/big"

	"github.com/3JoB/ulib/crypt/hmac"
)

const (
	Version string = "1.4.0"
)

type Session struct {
	client *Client
	d      *d

	Lotteryid string  // Lottery ID, if there is no one, you can call `NewLotteryID()` to generate one.
	UserNum   int     // Number of participants
	PrizeNum  int     // Quantity of prizes
	UserID    []int64 // All user IDs participating in the sweepstakes
}

type d struct {
	hashids   []string
	blockhash string
	seed      string
}

// Generate lottery seed
func (session *Session) seeds() {
	session.d.seed = hmac.SHA512(fmt.Sprintf("%v:%v:%v@%v", session.Lotteryid, session.UserNum, session.PrizeNum, session.d.blockhash), session.d.blockhash)
}

// Regenerate the lottery seed
func (session *Session) reSeed() {
	session.d.seed = hmac.SHA512(session.d.seed, session.d.blockhash)
}

// Get the latest block hash
func (session *Session) blockHash() error {
	header, err := session.client.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	session.d.blockhash = header.Hash().Hex()
	return nil
}

func (session *Session) getUser() []int64 {
	bigSeed, _ := new(big.Int).SetString(session.d.seed, 16)
	// var winner []int64
	winner := make([]int64, 0, session.PrizeNum)
	for i := 0; i < session.PrizeNum; i++ {
		winnerID := bigSeed.Mod(bigSeed, big.NewInt(int64(session.UserNum))).Int64()
		if winnerID != 0 {
			winner = append(winner, winnerID)
			session.reSeed()
			bigSeed, _ = new(big.Int).SetString(session.d.seed, 16)
		}
	}
	return winner
}

// Calculation draw results
func (session *Session) Do() ([]int64, error) {
	if session.Lotteryid == "" {
		session.NewLotteryID()
	}
	session.buildHash64()
	userlist := session.ids()
	if err := session.blockHash(); err != nil {
		return nil, err
	}

	session.seeds()
	winners := session.getUser()
	winnersID := make([]int, 0, len(winners))
	for _, winner := range winners {
		for i, userID := range userlist {
			if userID == winner {
				winnersID = append(winnersID, i)
				break
			}
		}
	}

	the_winners := make([]int64, 0, session.PrizeNum)
	for _, winnerID := range winnersID {
		if session.hash64(session.UserID[winnerID]) == session.d.hashids[winnerID] {
			the_winners = append(the_winners, session.UserID[winnerID])
		} else {
			for _, t := range session.UserID {
				if session.hash64(t) == session.d.hashids[winnerID] {
					the_winners = append(the_winners, session.UserID[winnerID])
					break
				}
			}
		}
	}

	if session.client.Debug {
		fmt.Printf("UserHashs: %v\nUserList: %v\n", session.d.hashids, userlist)
		fmt.Printf("BlockHash: %v\nSeed: %v\n", session.d.blockhash, session.d.seed)
		fmt.Printf("Winners: %v\nWinnersID: %v\nTheWinners: %v\n", winners, winnersID, the_winners)
	}

	return the_winners, nil
}
