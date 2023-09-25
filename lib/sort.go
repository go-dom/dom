//go:build !internel
// +build !internel

package dom

import "golang.org/x/exp/slices"

func sort[E []T, T string | int64](userIDs E) {
	slices.Sort(userIDs)
}

func sortSearch[E []T, T string | int64](userIDs E, key T) bool {
	_, ok := slices.BinarySearch(userIDs, key)
	return ok
}
