package dom

import (
	"context"
	"math/big"
	"sort"
	"strings"

	"github.com/3JoB/ulib/hash"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
)

type Session struct {
	client *Client
	d      *d

	Lotteryid string // Lottery ID, if there is no one, you can call `NewLotteryID()` to generate one.
	UserNum   int    // Number of participants
	PrizeNum  int
	Prize     []string // Quantity of prizes
	UserID    []int64  // All user IDs participating in the sweepstakes
}

type d struct {
	hashids   []string
	blockhash string
	seed      string
	winners   []int64
	winnernum int
}

// Generate lottery seed
func (session *Session) seeds() {
	data := litefmt.PSprint(
		session.Lotteryid,
		unsafeConvert.Itoa(session.PrizeNum),
		strings.Join(session.Prize, ","),
		session.d.blockhash,
	)
	session.d.seed = hash.SHA3_512S(data).Hex()
}

// Regenerate the lottery seed
func (session *Session) reSeed() {
	session.d.seed = hash.SHA3_512S(session.d.seed).Hex()
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

func (session *Session) getUser() {
	bigSeed, _ := new(big.Int).SetString(session.d.seed, 16)
	session.d.winners = make([]int64, 0, session.PrizeNum)
	for i := 0; i < session.PrizeNum; i++ {
		winnerID := bigSeed.Mod(bigSeed, big.NewInt(int64(session.UserNum))).Int64() + 1
		if session.winHas(winnerID - 1) {
			i--
		} else if winnerID != 0 {
			session.d.winners= append(session.d.winners, winnerID)
		}
		session.reSeed()
		bigSeed, _ = new(big.Int).SetString(session.d.seed, 16)
	}
	if session.isHas() {
		session.getUser()
	}
	session.d.winnernum = len(session.d.winners)
}

func (session *Session) isHas() bool {
	m := make(map[int64]int)
	for _, val := range session.d.winners {
		m[val]++
	}

	for _, value := range m {
		if value > 1 {
			return true
		}
	}
	return false
}

func (session *Session) winHas(id int64) bool {
	index := sort.Search(len(session.d.winners), func(i int) bool { return session.d.winners[i] >= id })
	if index < len(session.d.winners) && session.d.winners[index] == id {
		return true
	}
	return false
}
