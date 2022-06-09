var stack []int
var ans [][]int
var nn int
var kk int

func combine(n int, k int) [][]int {
    nn = n
    kk = k
    stack = []int{}
    ans = [][]int{}
    recursive(1)
    return ans
}

func recursive(i int) {
    if len(stack) > kk || len(stack) + nn - i + 1 < kk {
        return
    }
    if i == (nn + 1) {
        temp := make([]int, len(stack))
        copy(temp, stack)
        ans = append(ans, temp)
        return
    }
    
    recursive(i + 1)
    stack = append(stack, i)
    recursive(i + 1)
    stack = stack[:len(stack) - 1]
}