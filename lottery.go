package lottery

import (
	"context"
	"fmt"
	"math/big"

	"github.com/3JoB/ulib/crypt/hmac"
)

type Data struct {
	client *Client
	d      *d

	Lotteryid string
	UserNum   int // Number of participants
	PrizeNum  int // Quantity of prizes
	UserID    []int64
}

type d struct {
	hashids   []string
	blockhash string
	seed      string
}

func (stream *Data) seeds() {
	stream.d.seed = hmac.SHA512(fmt.Sprintf("%v:%v:%v@%v", stream.Lotteryid, stream.UserNum, stream.PrizeNum, stream.d.blockhash), stream.d.blockhash)
}

func (stream *Data) reSeed() {
	stream.d.seed = hmac.SHA512(stream.d.seed, stream.d.blockhash)
}

func (stream *Data) blockHash() error {
	header, err := stream.client.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	stream.d.blockhash = header.Hash().Hex()
	return nil
}

func (stream *Data) getUser() []int64 {
	bigSeed, _ := new(big.Int).SetString(stream.d.seed, 16)
	var winner []int64
	for i := 0; i < stream.PrizeNum; i++ {
		winnerID := bigSeed.Mod(bigSeed, big.NewInt(int64(stream.UserNum))).Int64()
		if winnerID != 0 {
			winner = append(winner, winnerID)
			stream.reSeed()
			bigSeed, _ = new(big.Int).SetString(stream.d.seed, 16)
		}
	}
	return winner
}

// Calculation draw results
func (stream *Data) Do() ([]int64, error) {
	stream.buildHash64()
	userlist := stream.ids()
	if err := stream.blockHash(); err != nil {
		return nil, err
	}

	var winnersID []int
	stream.seeds()
	winners := stream.getUser()
	for _, winner := range winners {
		for i, userID := range userlist {
			if userID == winner {
				winnersID = append(winnersID, i)
				break
			}
		}
	}

	the_winners := make([]int64, 0, stream.PrizeNum)
	for _, winnerID := range winnersID {
		if stream.hash64(stream.UserID[winnerID]) == stream.d.hashids[winnerID] {
			the_winners = append(the_winners, stream.UserID[winnerID])
		} else {
			for _, t := range stream.UserID {
				if stream.hash64(t) == stream.d.hashids[winnerID] {
					the_winners = append(the_winners, stream.UserID[winnerID])
					break
				}
			}
		}
	}

	if stream.client.Debug {
		fmt.Printf("UserHashs: %v\nUserList: %v\n", stream.d.hashids, userlist)
		fmt.Printf("BlockHash: %v\nSeed: %v\n", stream.d.blockhash, stream.d.seed)
		fmt.Printf("Winners: %v\nWinnersID: %v\nTheWinners: %v\n", winners, winnersID, the_winners)
	}

	return the_winners, nil
}
