func splitArray(nums []int, m int) int {
    left := 0
    right := 0
    for _, num := range nums {
        left = max(left, num)
        right += num
    }
    for left < right {
        mid := (left + right) / 2
        if validate(nums, m, mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return right
}

func validate(nums []int, m int, size int) bool {
    box := 0
    count := 1
    for _, num := range nums {
        if box + num <= size {
            box += num
        } else {
            count++
            box = num
        }
    }
    return count <= m
}

func max(v1, v2 int) int {
    if v1 > v2 {
        return v1
    } else {
        return v2
    }
}