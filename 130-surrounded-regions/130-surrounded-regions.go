var visited map[int]bool

func solve(board [][]byte) {
    x := len(board[0]) - 1
    y := len(board) - 1
    
    visited = make(map[int]bool, (x + 1) * (y + 1))
    for i := 0; i <= x; i++ {
        for j := 0; j <= y; j++ {
            if i == 0 || j == 0 || i == x || j == y {
                dfs(board, i, j)
            }
        }
    }
    
    for i := 0; i <= x; i++ {
        for j := 0; j <= y; j++ {
            if !visited[hash(i, j)] {
                board[j][i] = 88
            }
        }
    }
}

func dfs(board [][]byte, x, y int) {
    if x < 0 || x > len(board[0]) - 1 || y < 0 || y > len(board) - 1 || visited[hash(x, y)] {
        return
    }
    visited[hash(x, y)] = true
    if board[y][x] == 79 {
        dfs(board, x + 1, y)
        dfs(board, x - 1, y)
        dfs(board, x, y + 1)
        dfs(board, x, y - 1)
    }
}

func hash(x, y int) int {
    return y * 1000 + x
}
