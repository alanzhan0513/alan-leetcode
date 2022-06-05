func twoSum(nums []int, target int) []int {
    maps := make(map[int]int, len(nums))

    for i, num := range nums {
        if index, isExist := maps[num]; isExist {
            return []int{index, i}
        }
        maps[target - num] = i
    }
    return nil
}