func maxProduct(nums []int) int {
    length := len(nums)
    dpMax := make([]int, length)
    dpMin := make([]int, length)
    dpMax[0] = nums[0]
    dpMin[0] = nums[0]
    
    for i := 1; i < length; i++ {
        values := []int{dpMax[i - 1] * nums[i], dpMin[i - 1] * nums[i], nums[i]}
        dpMax[i] = max(values...)
        dpMin[i] = min(values...)
    }
    
    ans := dpMax[0]
    for i := 1; i < length; i++ {
        ans = max(ans, dpMax[i])
    }
    return ans
}

func max(v ...int) int {
    max := v[0]
    for i := 0; i < len(v); i++ {
        if v[i] > max {
            max = v[i]
        }
    }
    return max
}

func min(v ...int) int {
    min := v[0]
    for i := 0; i < len(v); i++ {
        if v[i] < min {
            min = v[i]
        }
    }
    return min
}