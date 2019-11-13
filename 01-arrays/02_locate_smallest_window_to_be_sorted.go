package arrays

import (
	"fmt"
)

// Given an array of integers that is out of order,
// determine the bounds of that smallest window that
// needs to be sorted in order of the entire array to
// be sorted.

func smallestSortedWindow(a []int) (int, int) {
	if len(a) == 0 {
		return -1, -1
	}
	start, end := len(a)-1, 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			start = min(start, i)
			end = max(i+1, end)
		} else if a[i] < a[start] {
			end = max(i, end)
		}
	}
	if a[len(a)-1] < a[start] {
		end = len(a) - 1
	}
	// check if result contains all values
	fmt.Printf("\t\tstart: %d, end: %d, input %v\n", start, end, a)
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
