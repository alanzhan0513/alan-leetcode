var to [][]int
var inDeg []int

func canFinish(numCourses int, prerequisites [][]int) bool {
    to = make([][]int, numCourses)
    inDeg = make([]int, numCourses)
    for _, prerequisite := range prerequisites {
        ai := prerequisite[0]
        bi := prerequisite[1]
        if to[bi] == nil {
            to[bi] = []int{ai}
        } else {
            to[bi] = append(to[bi], ai)
        }
        inDeg[ai]++
    }
    queue := []int{}
    // 拓墣排序，從 0 入度點出發
    for i := 0; i < numCourses; i++ {
        if inDeg[i] == 0 {
            queue = append(queue, i)
        }
    }
    lessons := []int{}
    for len(queue) != 0 {
        x := queue[0]
        queue = queue[1:len(queue)]
        lessons = append(lessons, x)
        // 擴展一個點，到周圍的點
        for _, y := range to[x] {
            inDeg[y]--
            // 入度為0 ，表示可以入隊
            if inDeg[y] == 0 {
                queue = append(queue, y)
            }
        }
    }
    return len(lessons) == numCourses
}