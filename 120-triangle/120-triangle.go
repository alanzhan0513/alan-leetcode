func minimumTotal(triangle [][]int) int {
    for i := len(triangle) - 2; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            if triangle[i + 1][j] > triangle[i + 1][j + 1] {
                triangle[i][j] += triangle[i + 1][j + 1]
            } else {
                triangle[i][j] += triangle[i + 1][j]
            }
        }
    }
    return triangle[0][0]
}