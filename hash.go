package lottery

import (
	"fmt"
	"sort"
	"time"

	"github.com/3JoB/ulib/crypt/hmac"
	"github.com/3JoB/ulid"
	"github.com/google/uuid"
	"lukechampine.com/frand"
)

func (stream *Data) newLotteryID() {
	id, _ := ulid.New(ulid.Timestamp(time.Now()), frand.New())
	ulids := ""
	if (id == ulid.ULID{}) {
		ulids = uuid.NewString()
	} else {
		ulids = id.String()
	}
	stream.Lotteryid = hmac.SHA512(fmt.Sprintf("%v$%v", uuid.NewString(), uuid.NewString()), ulids)
}

// Calculate user hash
func (stream *Data) hash64(userid int64) string {
	return hmac.SHA512(fmt.Sprintf("%v@%s", userid, stream.Lotteryid), stream.Lotteryid)
}

// Calculate user hash
func (stream *Data) buildHash64() {
	stream.d.hashids = make([]string, 0, len(stream.UserID))
	for _, userid := range stream.UserID {
		stream.d.hashids = append(stream.d.hashids, stream.hash64(userid))
	}
}

func (stream *Data) ids() []int64 {
	IDs := make(map[string]int)
	for i, userIDHash := range stream.d.hashids {
		IDs[userIDHash] = i
	}
	sort.Slice(stream.d.hashids, func(i, j int) bool {
		return stream.d.hashids[i] > stream.d.hashids[j]
	})
	userIDs := []int64{}
	for _, userIDHash := range stream.d.hashids {
		userIDs = append(userIDs, int64(IDs[userIDHash]))
	}
	return userIDs
}
