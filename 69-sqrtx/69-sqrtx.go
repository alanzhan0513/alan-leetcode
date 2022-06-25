func mySqrt(x int) int {
    left := 1
    right := x
    for left < right {
        mid := (left + right + 1) /2
        if (mid * mid <= x) {
            left = mid
        } else {
            right = mid -1
        }
    }
    return right
}