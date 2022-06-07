func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string, 0)
    for _, str := range strs {
        orderedStr := sortStringByCharacter(str)
        if v, isExist := groups[orderedStr]; isExist {
            groups[orderedStr] = append(v, str)
        } else {
            groups[orderedStr] = []string{str}
        }
    }
    ans := [][]string{}
    for _, v := range groups {
        ans = append(ans, v)
    }
    return ans
}

func stringToRuneSlice(v string) []rune {
    slice := []rune{}
    for _, r := range v {
        slice = append(slice, r)
    }
    return slice
}

func sortStringByCharacter(v string) string {
    r := stringToRuneSlice(v)
    
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    return string(r)
}