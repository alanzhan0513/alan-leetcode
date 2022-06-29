import "sort"
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func (i, j int) bool {
        return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
    })
    
    ans := make([][]int, 0)
    farthest := -1
    start := 0
    for _, interval := range intervals {
        left := interval[0]
        right := interval[1]
        if left <= farthest {
            farthest = max(farthest, right)
        } else {
            if farthest != -1 {
                ans = append(ans, []int{start, farthest})
            }
            start = left
            farthest = right
        }
    }
    ans = append(ans, []int{start, farthest})
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}