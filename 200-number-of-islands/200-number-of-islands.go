func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    
    m := len(grid)
    n := len(grid[0])
    ans := 0
    for x := 0; x < m; x++ {
        for y := 0; y < n; y++ {
            if grid[x][y] == '1' {
                ans += 1
                dfs(grid, x, y, m, n)
            }
        }
    }
    return ans
}

func dfs(grid [][]byte, x, y, m, n int) {
    if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] == '0' {
        return
    }
    grid[x][y] = '0'
    dfs(grid, x + 1, y, m, n)
    dfs(grid, x - 1, y, m, n)
    dfs(grid, x, y + 1, m, n)
    dfs(grid, x, y - 1, m, n)
}