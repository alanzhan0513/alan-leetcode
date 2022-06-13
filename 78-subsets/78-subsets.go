func subsets(nums []int) [][]int {
    return recursive(nums, 0, [][]int{[]int{}})
}

func recursive(nums []int, i int, ans [][]int) [][]int {
    if i >= len(nums) {
        return ans
    }
    
    temp := make([][]int, len(ans))
    for j := 0; j < len(temp); j++ {
        temp[j] = make([]int, len(ans[j]))
        copy(temp[j], ans[j])
        temp[j] = append(temp[j], nums[i])
    }
    
    ans = append(ans, temp...)
    return recursive(nums, i + 1, ans)
}