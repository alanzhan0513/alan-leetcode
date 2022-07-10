func lengthOfLIS(nums []int) int {
    length := len(nums)
    dp := make([]int, length)

    for i := 0; i < length; i++ {
        dp[i] = 1
    }
    
    for i:= 0; i < length; i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    ans := 0
    for i := 0; i < length; i++ {
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