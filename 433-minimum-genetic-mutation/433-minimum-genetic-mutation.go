func minMutation(start string, end string, bank []string) int {
    queue := []string{}
    visited := make(map[string]bool, 0)
    
    if contains(bank, end) && start == end {
        return -1
    }
    
    queue = append(queue, start)
    visited[start] = true
    ans := 0
    for len(queue) != 0 {
        length := len(queue)
        for i := 0; i < length; i++ {
            current := queue[0]
            queue = queue[1:]
            
            if current == end {
                return ans
            }
            
            for _, b := range bank {
                if canMutate(b, current) && !visited[b] {
                    queue = append(queue, b)
                    visited[b] = true
                }
            }
        }
        ans++
    }
    return -1
}

func canMutate(str1, str2 string) bool {
    counter := 0
    for i := 0; i < len(str1); i++ {
        if str1[i] != str2[i] {
            counter++
        }
    }
    return counter == 1
}

func contains(bank []string, tofind string) bool {
    for _, b := range bank {
        if tofind == b {
            return true
        }
    }
    return false
}
