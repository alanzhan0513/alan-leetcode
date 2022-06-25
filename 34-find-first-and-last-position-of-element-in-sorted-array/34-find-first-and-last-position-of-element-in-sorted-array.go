func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    
    ans := []int{}
    low := 0
    high := len(nums) - 1
    for low != high {
        mid := (low + high) / 2
        if nums[mid] >= target {
            high = mid
        } else {
            low = mid + 1
        }
    }
    if nums[high] != target {
        return []int{-1, -1} 
    }
    ans = append(ans, high)
    
    low = 0
    high = len(nums) - 1
    for low != high {
        mid := (low + high + 1) / 2
        if nums[mid] <= target {
            low = mid
        } else {
            high = mid - 1
        }
    }
    ans = append(ans, high)
    
    if ans[0] > ans[1] {
        return []int{-1, -1}
    }
    return []int{ans[0], ans[1]}
}