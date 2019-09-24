package arrays

//Given an array of integers, return a new array such that each element
//at index i of the new array is the product of all the numbers in the
//original array except one in i
func productOfTheOthers(nums []int) []int {
	product := 1
	for _, val := range nums {
		if val != 0 {
			product *= val
		}
	}
	result := make([]int, len(nums))
	for i, v := range nums {
		if v == 0 {
			result[i] = product
		} else {
			result[i] = product / v
		}
	}
	return result
}

func productOfTheOthersNoDivision(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = 1, 1
	for i := 1; i < n; i++ {
		if nums[i-1] != 0 {
			left[i] = left[i-1] * nums[i-1]
		} else {
			left[i] = left[i-1]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i+1] != 0 {
			right[i] = right[i+1] * nums[i+1]
		} else {
			right[i] = right[i+1]
		}
	}

	result := make([]int, n)
	for i := range nums {
		result[i] = left[i] * right[i]
	}
	return result
}
