func twoSum(nums []int, target int) []int {
    for i, a := range nums {
        for j, b := range nums {
            if i == j {
                continue
            }
            if a + b == target {
                return []int {i, j}
            }
        }
    }
    return nil
}