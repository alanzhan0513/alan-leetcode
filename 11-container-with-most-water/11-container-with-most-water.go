func maxArea(height []int) int {
    max := 0
    left := 0
    right := len(height) - 1

    for left != right {
        var area int
        if height[left] >= height[right] {
            area = height[right] * (right - left)
        } else {
            area = height[left] * (right - left)
        }
        
        if area > max {
            max = area
        }
        
        if height[left] >= height[right] {
            right--
        } else {
            left++
        }
    }
    
    return max
}