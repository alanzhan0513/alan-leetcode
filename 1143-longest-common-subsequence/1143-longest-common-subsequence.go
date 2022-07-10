func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
    n := len(text2)
    text1 = " " + text1
    text2 = " " + text2
    dpGrid := make([][]int, m + 1)
    for i := 0; i <= m; i++ {
        dpGrid[i] = make([]int, n + 1)
    }
    for i := 1; i <= m ; i++ {
        for j := 1; j <= n; j++ {
            if text1[i] == text2[j]{
                dpGrid[i][j] = dpGrid[i - 1][j - 1] + 1
            } else {
                dpGrid[i][j] = max(dpGrid[i - 1][j], dpGrid[i][j - 1])
            }
        }
    }
    return dpGrid[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}