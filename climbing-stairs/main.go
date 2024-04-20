package main

import "fmt"

func main() {
	fmt.Println(climbStairsV2(2))
}

func climbStairs(n int) int {
	one := 1
	two := 1
	var temp int
	for range n - 1 {
		temp = one
		one = one + two
		two = temp
	}
	return one
}

func climbStairsV2(n int) int {
	if n <= 1 {
		return n
	}
	res := []int{1, 2}

	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}

	return res[len(res)-1]

}
