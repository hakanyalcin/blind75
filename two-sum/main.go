package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil
}

func twoSumV3(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		x := target - num
		if j, ok := m[x]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func practiceA(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		val, ok := m[target-num]
		if ok {
			return []int{val, i}
		}
		m[num] = i
	}
	return nil
}
