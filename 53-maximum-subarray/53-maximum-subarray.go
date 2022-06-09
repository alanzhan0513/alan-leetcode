func maxSubArray(nums []int) int {
    max := nums[0]
    tmp := nums[0]
    for i := 1; i < len(nums); i++ {
        if tmp > 0 {
            tmp += nums[i]
        } else {
            tmp = nums[i]
        }
        
        if tmp > max {
            max = tmp
        }
    }
    return max
}