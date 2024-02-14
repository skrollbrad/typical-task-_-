package main

import "fmt"

func maxProfit(prices []int) int {

	minPrice := prices[0]

	profit := 0

	for i := 1; i < len(prices); i++ {
		if minPrice < prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit

}
func main() {

	prices := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices))

}
