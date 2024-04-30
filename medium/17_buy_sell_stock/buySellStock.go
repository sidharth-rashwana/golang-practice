// Best Time to Buy and Sell Stock
package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		// update the minimum price seen so far
		if price < minPrice {
			minPrice = price
		}

		// calculate profit if sold today
		profit := price - minPrice

		// update max profit
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // Output: 5
}
