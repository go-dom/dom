package dom

import (
	"context"
	"fmt"
	"math/big"

	errs "github.com/3JoB/ulib/err"
	"github.com/3JoB/ulib/hash/hmac"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
)

const (
	Version string = "1.5.0"
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
	session.d.seed = hmac.SHA3_512S(litefmt.Sprint(session.Lotteryid, ":", unsafeConvert.IntToString(session.UserNum), ":", unsafeConvert.IntToString(session.PrizeNum), "@", session.d.blockhash), session.d.blockhash).Hex()
}

// Regenerate the lottery seed
func (session *Session) reSeed() {
	session.d.seed = hmac.SHA3_512S(session.d.seed, session.d.blockhash).Hex()
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

var ErrLess error = &errs.Err{Op: "dom", Err: "The number of prizes cannot be less than the number of participants!"}

// Calculation draw results
func (session *Session) Do() ([]int64, error) {
	if session.Lotteryid == "" {
		session.NewLotteryID()
	}
	if session.PrizeNum >= session.UserNum {
		return nil, ErrLess
	}
	session.buildHash64()
	userlist := session.ids()
	if err := session.blockHash(); err != nil {
		return nil, err
	}

	session.seeds()
	winners := session.getUser()
	winum := len(winners)
	if winum == 0 {
		winners = session.getUser()
	} else if winum < session.PrizeNum {
		winners = session.getUser()
	}
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
		if session.hash(session.UserID[winnerID]) == session.d.hashids[winnerID] {
			the_winners = append(the_winners, session.UserID[winnerID])
		} else {
			for _, t := range session.UserID {
				if session.hash(t) == session.d.hashids[winnerID] {
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
