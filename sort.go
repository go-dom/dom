//go:build !internel
// +build !internel

package dom

import "golang.org/x/exp/slices"

func sort[E []T, T string | int64](datas E) {
	slices.Sort(datas)
}

func sortSearch[E []T, T string | int64](datas E, key T) bool {
	/*for _, v := range datas {
		if v == key {
			return true
		}
	}
	return false*/
	sort(datas)
	_, ok := slices.BinarySearch(datas, key)
	return ok
}
