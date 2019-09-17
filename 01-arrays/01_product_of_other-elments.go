package arrays

//Given an array of integers, return a new array such that each element
//at index i of the new array is the product of all the numbers in the
//original array except one in i
func Product(nums []int) []int {
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
