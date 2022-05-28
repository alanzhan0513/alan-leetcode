func moveZeroes(nums []int)  {
    n := 0
    numsLen := len(nums)
    for i := 0; i < numsLen; i++ {
        if nums[i] != 0 {
            nums[n] = nums[i]
            n++
        }
    }
    for n < numsLen {
        nums[n] = 0
        n++
    }
}