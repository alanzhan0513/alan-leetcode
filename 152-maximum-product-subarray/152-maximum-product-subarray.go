func maxProduct(nums []int) int {
    max := nums[0]
    min := nums[0]
    ans := max
    
    for i := 1; i < len(nums); i++ {
        n := nums[i]
        tempMax := getMax(n, max * n, min * n)
        min = getMin(n, max * n, min * n)
        
        max = tempMax
        ans = getMax(max, ans)
    }
    
    return ans
}

func getMax(nums ...int) int {
    max := nums[0]
    for _, n := range nums {
        if n > max {
            max = n
        }
    }
    return max
}

func getMin(nums ...int) int {
    min := nums[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
    }
    return min
}