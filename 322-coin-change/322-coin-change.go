func coinChange(coins []int, amount int) int {
    maxInt := 1 << 31 - 1
    dp := make([]int, amount + 1)
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        dp[i] = maxInt
        for j := 0; j < len(coins); j++ {
            if i - coins[j] >= 0 {
                dp[i] = min(dp[i], dp[i - coins[j]] + 1)
            }
        }
    }
    if dp[amount] == maxInt {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}