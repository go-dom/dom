package dom

import (
	"fmt"

	errs "github.com/3JoB/ulib/err"
)

type SessionInfo struct {
	Block     string        `json:"block_hash"` // Block Hash
	Lotteryid string        `json:"lottery_id"` // Lottery ID, if there is no one, you can call `NewLotteryID()` to generate one.
	UserNum   int           `json:"user_num"`   // Number of participants
	PrizeNum  int           `json:"prize_num"`  // Quantity of prizes
	UserID    []UserResults `json:"joined"`     // All user IDs participating in the sweepstakes
	Results   []UserResults `json:"results"`    // Winners
}

type UserResults struct {
	UserID   int64  `json:"user_id"`
	UserHash string `json:"user_hash"`
}

const Version string = "1.6.0"

var ErrDataEmpty error = &errs.Err{Op: "dom", Err: "data can not be empty!"}
var ErrLess error = &errs.Err{Op: "dom", Err: "The number of prizes cannot be less than the number of participants!"}
var ErrMoreRetry error = &errs.Err{Op: "dom", Err: "Unable to correctly generate winnerid, and too many retries, this operation has been terminated."}
var ErrPrizeNum error = &errs.Err{Op: "dom", Err: "TIncorrect number of prizes"}
var ErrUserNum error = &errs.Err{Op: "dom", Err: "Incorrect number of users"}

// Calculation draw results
func (session *Session) Do() (SessionInfo, error) {
	sessionInfo := SessionInfo{}

	if session.Lotteryid == "" {
		session.NewLotteryID()
	}
	if session.PrizeNum >= session.UserNum {
		return sessionInfo, ErrLess
	}
	if session.PrizeNum < 1 {
		return sessionInfo, ErrPrizeNum
	}
	if session.UserNum < 1 {
		return sessionInfo, ErrUserNum
	}
	session.buildHash64()
	userlist := session.ids()
	if err := session.blockHash(); err != nil {
		return sessionInfo, err
	}
	sessionInfo.Block = session.d.blockhash

	session.seeds()
	session.getUser()
	retry := 0

Retrys:
	if retry > 4 {
		if session.client.Debug {
			fmt.Printf("UserHashs: %v\nUserList: %v\n", session.d.hashids, userlist)
			fmt.Printf("BlockHash: %v\nSeed: %v\n", session.d.blockhash, session.d.seed)
		}
		return sessionInfo, ErrMoreRetry
	}

	if session.d.winnernum == 0 {
		session.getUser()
		retry++
		goto Retrys
	} else if session.d.winnernum != session.PrizeNum {
		session.getUser()
		retry++
		goto Retrys
	}

	retry = 0

	winnersID := make([]int, 0, session.d.winnernum)
	for _, winner := range session.d.winners {
		for i, userID := range userlist {
			if userID == winner {
				winnersID = append(winnersID, i)
				break
			}
		}
	}

	sessionInfo.Results = make([]UserResults, 0, session.PrizeNum)
	for _, winnerID := range winnersID {
		if session.hash(session.UserID[winnerID]) == session.d.hashids[winnerID] {
			sessionInfo.Results = append(sessionInfo.Results, UserResults{UserID: session.UserID[winnerID], UserHash: session.d.hashids[winnerID]})
		} else {
			for _, t := range session.UserID {
				if session.hash(t) == session.d.hashids[winnerID] {
					sessionInfo.Results = append(sessionInfo.Results, UserResults{UserID: session.UserID[winnerID], UserHash: session.d.hashids[winnerID]})
					break
				}
			}
		}
	}

	if session.client.Debug {
		fmt.Printf("UserHashs: %v\nUserList: %v\n", session.d.hashids, userlist)
		fmt.Printf("BlockHash: %v\nSeed: %v\n", session.d.blockhash, session.d.seed)
		fmt.Printf("Winners: %v\nWinnersID: %v\nTheWinners: %v\n", session.d.winners, winnersID, sessionInfo.Results)
	}

	return sessionInfo, nil
}
