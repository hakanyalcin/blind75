package main

import "fmt"

func main() {

	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}

// func missingNumber(nums []int) int {

// 	allNums := make([]int, len(nums)+1)
// 	for i := 0; i <= len(nums); i++ {
// 		allNums[i] = i
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		allNums = append(allNums[:nums[i]], allNums[nums[i]+1:]...)
// 	}
// 	return 0
// }

func missingNumber(nums []int) int {
	var n int = len(nums)

	expected_sum := (n * (n + 1)) / 2
	calculated_sum := 0

	for i := 0; i < n; i++ {
		calculated_sum += nums[i]
	}
	return expected_sum - calculated_sum
}
