//go:build internel
// +build internel

package dom

import "gopkg.in/dom.v2/slices"

func sort[E []T, T string | int64](datas E) {
	slices.Sort(datas)
}

func sortSearch[E []T, T string | int64](datas E, key T) bool {
	sort(datas)
	_, ok := slices.BinarySearch(datas, key)
	return ok
}
