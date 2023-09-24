//go:build !internel
// +build !internel

package dom

import "gopkg.in/dom.v2/slices"

func sortUserIDs[E []T, T string | int64](userIDs E) {
	slices.Sort(userIDs)
}

func sortSearch[E []T, T string | int64](userIDs E, key T) bool {
	_, ok := slices.BinarySearch(userIDs, key)
	return ok
}
