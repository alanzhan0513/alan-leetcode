type pointer struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {
    set := make(map[pointer]bool, 0)
    for _, obstacle := range obstacles {
        set[pointer{obstacle[0], obstacle[1]}] = true
    }
    
    x, y := 0, 0
    direction := 0 // n 0, e 1, s 2, w 3
    dx := [4]int{0, 1, 0, -1}
    dy := [4]int{1, 0, -1, 0}
    ans := 0
    for _, command := range commands {
        if command == -1 {
            // (direction + 1) % 4 turn right
            direction = (direction + 1) % 4
        } else if command == -2 {
            // (direction - 1 + 4) % 4 turn left
            direction = (direction + 3) % 4
        } else {
            for i := 0; i < command; i++ {
                nextX := x + dx[direction]
                nextY := y + dy[direction]
                
                if _, isExist := set[pointer{nextX, nextY}]; isExist {
                     break
                }
                x = nextX
                y = nextY
                
                if ans < x * x + y * y {
                    ans = x * x + y * y
                }
            }
        }
    }
    return ans
}

