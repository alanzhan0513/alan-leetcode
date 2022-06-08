type NumMatrix struct {
    Dp [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    numMatrix := NumMatrix{make([][]int, len(matrix) + 1)}
    for i := range numMatrix.Dp {
        numMatrix.Dp[i] = make([]int, len(matrix[0]) + 1)
    }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            numMatrix.Dp[i + 1][j + 1] = numMatrix.Dp[i + 1][j] + numMatrix.Dp[i][j + 1] + matrix[i][j] - numMatrix.Dp[i][j];
        }
    }
    return numMatrix
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.Dp[row2 + 1][col2 + 1] - this.Dp[row1][col2 + 1] - this.Dp[row2 + 1][col1] + this.Dp[row1][col1];
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */