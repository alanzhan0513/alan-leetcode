func maxSubArray(nums []int) int {
    length := len(nums)
    dp := make([]int, length)
    dp[0] = nums[0]
    for i := 1; i < length; i++ {
        dp[i] = max(dp[i - 1] + nums[i], nums[i])
    }
    ans := dp[0]
    for i := range dp {
        ans = max(ans, dp[i])
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}