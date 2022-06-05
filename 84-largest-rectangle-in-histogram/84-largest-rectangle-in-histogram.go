type rect struct {
    width int
    height int
}

func largestRectangleArea(heights []int) int {
    stack := []rect{}
    answer := 0
    heights = append(heights, 0)
    for _, height := range heights {
        accumulatedWidth := 0
        
        for len(stack) != 0 && stack[len(stack) - 1].height > height {
            accumulatedWidth += stack[len(stack) - 1].width
            temp := stack[len(stack) - 1].height * accumulatedWidth
            if temp > answer {
                answer = temp
            }
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, rect{accumulatedWidth + 1, height})
    }
    return answer
}
