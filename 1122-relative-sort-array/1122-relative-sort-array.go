func relativeSortArray(arr1 []int, arr2 []int) []int {
    count := make([]int, 1001)
    for _, v := range arr1 {
        count[v]++
    }
    ans := []int{}
    for _, v := range arr2 {
        for count[v] > 0 {
            ans = append(ans, v)
            count[v]--
        }
    }
    for i := 0; i <= 1000; i++ {
        for count[i] > 0 {
            ans = append(ans, i)
            count[i]--
        }
    }
    return ans
}