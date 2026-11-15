package besttimetobuyandsellstock

import (
	"math"
)

func MaxProfitBruteForce(prices []int) int {
	var maxProfit int

	for i := range prices {
		for j := i + 1; j < len(prices); j++ {
			maxProfit = max(maxProfit, prices[j]-prices[i])
		}
	}

	return maxProfit
}

func MaxProfit(prices []int) int {
	var maxProfit int
	minPrice := math.MaxInt

	for _, price := range prices {
		minPrice = min(minPrice, price)

		dayProfit := price - minPrice
		maxProfit = max(maxProfit, dayProfit)
	}

	return maxProfit
}
