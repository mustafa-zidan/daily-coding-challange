package arrays

import "daily-coding-chalange/utils"

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

//what if the elements can wrap around? for example [8, -1, 3, 4] returns 15
func maxCircularSubarraySum(a []int) int {
	maxSum := sum(a) - minSubarraySum(a)
	return utils.Max(maxSum, maxSubarraySum(a))
}

func minSubarraySum(a []int) int {
	min, minSoFar := 0, 0

	for _, v := range a {
		minSoFar = utils.Min(v, minSoFar+v)
		min = utils.Min(min, minSoFar)
	}
	return min
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
