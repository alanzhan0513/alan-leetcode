func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    
    if n < 0 {
        return 1 / myPow(x, -n)
    } 
    
    temp := myPow(x, n / 2)
    ans := temp * temp
    if (n % 2 == 1) {
        ans *= x
    }
    return ans
}