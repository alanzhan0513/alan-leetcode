func maxProfit(prices []int) int {
    maxprofit := 0
    for i := 1; i < len(prices); i++ {
        maxprofit += max(0, prices[i] - prices[i - 1])
    }
    return maxprofit
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}