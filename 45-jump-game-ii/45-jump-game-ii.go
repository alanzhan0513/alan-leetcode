func jump(nums []int) int {
    now := 0
    ans := 0
    for now < len(nums) - 1 {
        right := now + nums[now]
        if right >= len(nums) - 1 {
            return ans + 1
        }
        nextRight := right
        next := now
        for i := now + 1; i <= right; i++ {
            for i + nums[i] > nextRight {
                nextRight = i + nums[i]
                next = i
            }
        }
        now = next
        ans++
    }
    return ans
}