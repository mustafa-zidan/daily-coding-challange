package arrays

import (
	"math"
)

// Given an array of integers that is out of order,
// determine the bounds of that smallest window that
// needs to be sorted in order of the entire array to
// be sorted.

func smallestSortedWindow(a []int) (int, int) {
	if len(a) == 0 {
		return -1, -1
	}
	minSeen, maxSeen := math.MinInt32, math.MaxInt32
	start, end := len(a)-1, 0
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			start = min(start, i-1)
			minSeen = a[start]
		} else if a[i] < minSeen || a[i] < maxSeen {
			end = max(i, end)
			maxSeen = a[end]
		}
	}
	// check if result contains all values
	if start > end {
		return -1, -1
	}
	return start, end
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
