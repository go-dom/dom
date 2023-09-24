//go:build !internel
// +build !internel

package dom

import "golang.org/x/exp/slices"

func sortUserIDs(userIDs []string) {
	slices.Sort(userIDs)
}
