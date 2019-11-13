package arrays

import "github.com/mustafa-zidan/daily-coding-chalange/utils"

// Given an array of numbers, find the maximum
// sum of any contiguous subarray of the array,

func maxSubarraySum(a []int) int {
	max, maxSoFar := 0, 0

	for _, v := range a {
		maxSoFar = utils.Max(v, maxSoFar+v)
		max = utils.Max(max, maxSoFar)
	}
	return max
}
