package dom

import (
	"sort"
	"time"

	"github.com/3JoB/ulib/hash/hmac"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/ulid"
	"github.com/3JoB/unsafeConvert"
	"github.com/google/uuid"
	"lukechampine.com/frand"
)

func (session *Session) newLotteryID() {
	id, _ := ulid.New(ulid.Timestamp(time.Now()), frand.New())
	ulids := ""
	if (id == ulid.ULID{}) {
		ulids = uuid.NewString()
	} else {
		ulids = id.String()
	}
	session.Lotteryid = hmac.SHA3_512S(litefmt.Sprint(uuid.NewString(), "$", uuid.NewString()), ulids).Hex()
}

// Calculate user hash
func (session *Session) hash(userid int64) string {
	return hmac.SHA3_512S(litefmt.Sprint(unsafeConvert.Int64ToString(userid), "@", session.Lotteryid), session.Lotteryid).Hex()
}

// Calculate user hash
func (session *Session) buildHash64() {
	session.d.hashids = make([]string, 0, len(session.UserID))
	for _, userid := range session.UserID {
		session.d.hashids = append(session.d.hashids, session.hash(userid))
	}
}

func (session *Session) ids() []int64 {
	IDs := make(map[string]int, len(session.d.hashids))
	for i, userIDHash := range session.d.hashids {
		IDs[userIDHash] = i
	}
	sort.Slice(session.d.hashids, func(i, j int) bool {
		return session.d.hashids[i] > session.d.hashids[j]
	})
	userIDs := make([]int64, 0, len(session.d.hashids))
	// userIDs := []int64{}
	for _, userIDHash := range session.d.hashids {
		userIDs = append(userIDs, int64(IDs[userIDHash]))
	}
	return userIDs
}
