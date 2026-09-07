func maxProfit(prices []int) int {
    min := prices[0]
    profit := 0
    for _, price := range prices {
        if price < min {
            min = price
        }
        if price - min > profit {
            profit = price - min
        }
    }
    return profit
}
