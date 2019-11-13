package arrays

import "github.com/mustafa-zidan/daily-coding-chalange/utils"

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
			start = utils.Min(start, i)
			end = utils.Max(i+1, end)
		} else if a[i] < a[start] {
			end = utils.Max(i, end)
		}
	}
	if a[len(a)-1] < a[start] {
		end = len(a) - 1
	}
	// check if result contains all values
	if start > end {
		return -1, -1
	}
	return start, end
}
