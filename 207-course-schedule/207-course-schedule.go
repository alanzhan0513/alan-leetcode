
func canFinish(numCourses int, prerequisites [][]int) bool {
    to := make([][]int, numCourses)
    inDeg := make([]int, numCourses)
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
        for _, y := range to[x] {
            inDeg[y]--
            if inDeg[y] == 0 {
                queue = append(queue, y)
            }
        }
    }
    return len(lessons) == numCourses
}