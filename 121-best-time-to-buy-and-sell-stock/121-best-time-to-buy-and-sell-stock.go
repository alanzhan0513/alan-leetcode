func maxProfit(prices []int) int {
    min := prices[0]
    profit := 0
    for i := 1; i < len(prices); i++ {
        if min > prices[i] {
            min = prices[i]
        } else {
            profit = max(profit, prices[i] - min)
        }
    }
    return profit
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}