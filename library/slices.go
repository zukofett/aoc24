package aoc

import "slices"

func SlicesIntersect[S ~[]E, E comparable](first S, second S) bool {
	for _, item := range second {
		if slices.Contains(first, item) {
			return true
		}
	}
	return false
}
