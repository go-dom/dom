package lottery

import (
	"fmt"
	"sort"

	"github.com/3JoB/ulib/crypt/hmac"
)

func Hash(userid, lotteryid string) string {
	return hmac.SHA512(fmt.Sprintf("%s:%s", userid, lotteryid), lotteryid)
}

func BuildHashs(userid []string, lotteryid string) []string {
	hashes := make([]string, 0, len(userid))
	for _, userid := range userid {
		hashes = append(hashes, Hash(userid, lotteryid))
	}
	return hashes
}

func Hash64(userid int64, lotteryid string) string {
	return hmac.SHA512(fmt.Sprintf("%s:%s", userid, lotteryid), lotteryid)
}

func BuildHash64(userid []int64, lotteryid string) []string {
	hashes := make([]string, 0, len(userid))
	for _, userid := range userid {
		hashes = append(hashes, Hash64(userid, lotteryid))
	}
	return hashes
}

func IDS(hashsID []string) []int {
	IDs := make(map[string]int)
	for i, userIDHash := range hashsID {
		IDs[userIDHash] = i
	}
	sort.Slice(hashsID, func(i, j int) bool {
		return hashsID[i] > hashsID[j]
	})
	userIDs := []int{}
	for _, userIDHash := range hashsID {
		userIDs = append(userIDs, IDs[userIDHash])
	}
	return userIDs
}
