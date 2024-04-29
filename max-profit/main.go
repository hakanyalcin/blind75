package main

import "fmt"

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{7, 6, 4, 3, 1}
	prices := []int{2, 1, 2, 0, 1}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if updatedProfit := prices[i] - minPrice; updatedProfit > profit {
			profit = updatedProfit
		}
	}
	return profit
}
