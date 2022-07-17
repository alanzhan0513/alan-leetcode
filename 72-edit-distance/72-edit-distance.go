func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(word2) + 1)
        dp[i][0] = i
    }
    
    for j := 0; j < len(word2) + 1; j++ {
        dp[0][j] = j
    }
    
    for i := 1; i < len(word1) + 1; i++ {
        for j := 1; j < len(word2) + 1; j++ {
            if word1[i - 1] == word2[j - 1] {
                dp[i][j] = 1 + min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1] - 1)
            } else {
                dp[i][j] = 1 + min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1])
            }
        }
    }
    return dp[len(word1)][len(word2)]
}

func min(values... int) int {
    min := 1 << 32 - 1
    for _, v := range values {
        if v < min {
            min = v
        }
    }
    return min
}