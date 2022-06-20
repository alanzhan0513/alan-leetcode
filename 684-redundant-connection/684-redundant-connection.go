var to [][]int
var visited []bool
var hasCycle bool
func findRedundantConnection(edges [][]int) []int {
    n := 0
    for _, edge := range edges {
        x := edge[0]
        y := edge[1]
        values := []int{n, x, y}
        n = max(values...)
    }
    to = make([][]int, n + 1)
    visited = make([]bool, n + 1)
    for _, edge := range edges {
        x := edge[0]
        y := edge[1]
        if to[x] == nil {
            to[x] = []int{y}
        } else {
            to[x] = append(to[x], y)
        }
        if to[y] == nil {
            to[y] = []int{x}
        } else {
            to[y] = append(to[y], x)
        }
        hasCycle = false
        for i := 0; i <= n; i++ {
            visited[i] = false
        }
        dfs(x, 0)
        if hasCycle {
            return edge
        }
    }
    return []int{}
}

func dfs(x int, father int) {
    visited[x] = true
    for _, y := range to[x] {
        if y == father {
            continue
        }
        if !visited[y] {
            dfs(y, x)
        } else {
            hasCycle = true
        }
    }
}

func max(values ...int) int {
    max := 0
    for _, v := range values {
        if v > max {
            max = v
        }
    }
    return max
}