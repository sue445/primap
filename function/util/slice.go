package util

import (
	"github.com/deckarep/golang-set"
	"sort"
)

// SubtractSlice returns subtracted slice (src - sub)
func SubtractSlice(src []string, sub []string) []string {
	return subtractSliceWithSet(src, sub)
}

func subtractSliceWithContains(src []string, sub []string) []string {
	var ret []string

	for _, s := range src {
		if !Contains(sub, s) {
			ret = append(ret, s)
		}
	}

	return ret
}

// Contains returns whether slice contains item
func Contains(slice []string, item string) bool {
	// c.f. https://stackoverflow.com/questions/10485743/contains-method-for-a-slice

	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func subtractSliceWithSet(src []string, sub []string) []string {
	subSet := mapset.NewSet()

	for _, s := range sub {
		subSet.Add(s)
	}

	var ret []string

	for _, s := range src {
		if !subSet.Contains(s) {
			ret = append(ret, s)
		}
	}

	return ret
}

// SortedSlice returns sorted slice
func SortedSlice(src []string) []string {
	if sort.StringsAreSorted(src) {
		return src
	}

	var dst []string
	for _, s := range src {
		dst = append(dst, s)
	}
	sort.Strings(dst)

	return dst
}

// UniqueSlice returns slice with duplicate elements eliminated
func UniqueSlice(src []string) []string {
	set := mapset.NewSet()

	var ret []string

	for _, s := range src {
		if !set.Contains(s) {
			ret = append(ret, s)
			set.Add(s)
		}
	}

	return ret

}
