package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	i, j := 0, 0
	res := []int{}
	for i = 0; i < len(nums); i++ {
		for j = i + i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, nums[i])
				res = append(res, nums[j])
				return res
			}
		}
	}
	return res
}
