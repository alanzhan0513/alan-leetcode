type rect struct {
    width   int
    height  int
}

func largestRectangleArea(heights []int) int {
    stack := []rect{}
    max := 0
    heights = append(heights, 0)
    for _, height := range heights {
        accumulatedWidth := 0
        
        for len(stack) != 0 && stack[len(stack) - 1].height >= height {
            top := stack[len(stack) - 1]
            accumulatedWidth += top.width
            if accumulatedWidth * top.height > max {
                max = accumulatedWidth * top.height
            }
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, rect{accumulatedWidth + 1, height})
    }
    
    return max
}