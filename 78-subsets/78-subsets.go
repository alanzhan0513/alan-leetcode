var length int
var stack []int
var ans [][]int

func subsets(nums []int) [][]int {
    length = len(nums)
    stack = []int{}
    ans = [][]int{}
    recursive(0, nums)
    return ans
}

func recursive(i int, nums []int) {
    if length == i {
        temp := make([]int, len(stack))
        copy(temp, stack)
        ans = append(ans, temp)
        return
    }
    
    recursive(i + 1, nums)
    stack = append(stack, nums[i])
    recursive(i + 1, nums)
    stack = stack[:len(stack) - 1]
}