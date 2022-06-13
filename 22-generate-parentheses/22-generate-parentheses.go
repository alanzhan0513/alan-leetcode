func generateParenthesis(n int) []string {
    if n == 0 {
        return []string{""}
    }
    ans := []string{}
    for i := 1; i <= n; i++ {
        A := generateParenthesis(i - 1)
        B := generateParenthesis(n - i)
        
        for _, a := range A {
            for _, b := range B {
                ans = append(ans, "(" + a + ")" + b)
            }
        }
    }
    return ans
}