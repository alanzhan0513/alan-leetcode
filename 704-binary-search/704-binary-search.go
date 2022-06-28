func search(nums []int, target int) int {
    right := len(nums) - 1
    if right == -1 {
        return -1
    }
    left := 0
    
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}