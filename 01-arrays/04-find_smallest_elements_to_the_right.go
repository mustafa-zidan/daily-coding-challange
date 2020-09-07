package arrays

// Given an array of integers, return a new
// array where each element in the new array
// is the number of smaller element in the
// original input array

// Time O(N logN) Space O(N)
func smallestCount(a []int) []int {
	seen := make([]int, 1)
	seen[0], a[len(a)-1] = a[len(a)-1], 0
	for i := len(a) - 2; i >= 0; i-- {
		index := getElementIndex(seen, a[i])
		seen = append(seen[:index], append([]int{a[i]}, seen[index:]...)...)
		a[i] = index
	}
	return a
}

// Return the element possible index in the array
func getElementIndex(a []int, e int) int {
	if len(a) == 0 {
		return 0
	}
	min, max, pivot := 0, len(a), 0

	for min < max {
		if a[pivot] > e {
			max = pivot - 1
		} else {
			min = pivot + 1
		}
		pivot = (min + max) / 2
	}
	return pivot
}

// Time O(n^2) Space O(1)
func smallestCountBrute(a []int) []int {
	for i := 0; i < len(a); i++ {
		count := 0
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				count++
			}
		}
		a[i] = count
	}
	return a
}
