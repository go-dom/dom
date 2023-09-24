//go:build internel
// +build internel

package dom

import "gopkg.in/dom.v2/slices"

func sortUserIDs(userIDs []string) {
	slices.Sort(userIDs)
}
