func mySqrt(x int) int {
    left := 1
    right := x
    for left < right {
        mid := (left + right + 1) / 2
        if mid * mid > x {
            right = mid - 1
        } else {
            left = mid
        }
    }
    return right
}