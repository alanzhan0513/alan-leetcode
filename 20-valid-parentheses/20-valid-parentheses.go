func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    stack := make([]rune, 0)
    for _, v := range s {
        if v == '(' || v == '[' || v == '{' {
            stack = append(stack, v)
        } else if len(stack) > 0 &&
        ((v == ')' && stack[len(stack) - 1] == '(') ||
        (v == ']' && stack[len(stack) - 1] == '[')||
         (v == '}' && stack[len(stack) - 1] == '{')) {
            stack = stack[:len(stack) - 1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}