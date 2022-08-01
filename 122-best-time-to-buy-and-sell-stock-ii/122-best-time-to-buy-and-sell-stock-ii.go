func maxProfit(prices []int) int {
    // move index to 1
    prices = append([]int{ 0 }, prices...)
    
    dp := make([][]int, len(prices))
    dp[0] = make([]int, 2)
    dp[0][1] = -1 << 32
    dp[0][0] = 0
    
    for i := 1; i < len(prices); i++ {
        dp[i] = make([]int, 2)
        dp[i][0] = -1 << 32
        dp[i][1] = -1 << 32
        
        dp[i][1] = max(dp[i][1], dp[i - 1][0] - prices[i])
        dp[i][0] = max(dp[i][0], dp[i - 1][1] + prices[i])
        for j := 0; j < 2; j++ {
            dp[i][j] = max(dp[i][j], dp[i - 1][j])
        }
    }
    return dp[len(prices) - 1][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}