var maps map[rune][]rune
func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    
    maps = make(map[rune][]rune, 8)
    maps['2'] = []rune{'a','b','c'}
    maps['3'] = []rune{'d','e','f'}
    maps['4'] = []rune{'g','h','i'}
    maps['5'] = []rune{'j','k','l'}
    maps['6'] = []rune{'m','n','o'}
    maps['7'] = []rune{'p','q','r','s'}
    maps['8'] = []rune{'t','u','v'}
    maps['9'] = []rune{'w','x','y','z'}

    ans := []string{}
    recursive(digits, 0, "", &ans)
    return ans
}

func recursive(digits string, index int, current string, ans *[]string) {
    if index == len(digits) {
        *ans = append(*ans, current)
        return
    }

    for _, v := range maps[rune(digits[index])] {
        current += string(v)
        recursive(digits, index + 1, current, ans)
        current = current[:len(current) - 1]
    }
}