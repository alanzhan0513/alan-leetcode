func rob(nums []int) int {
    length := len(nums)
    dp := make([]int, length)
    dp[0] = nums[0]

    for i := 1; i < length; i++ {
        if i < 2 {
            dp[i] = max(nums[i], dp[i - 1])
        } else {
            dp[i] = max(nums[i] + dp[i - 2], dp[i - 1])
        }
    }
    
    return dp[length - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}