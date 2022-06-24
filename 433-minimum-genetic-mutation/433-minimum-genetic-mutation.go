func minMutation(start string, end string, bank []string) int {
	mp := make(map[string]bool)
	vis := make(map[string]bool)
	for _, v := range bank {
    
		mp[v] = true
	}
	ans := math.MaxInt32
	dirs := []byte("ACGT")
	var dfs func(cur string, step int)
	dfs = func(cur string, step int) {
    
		if cur == end {
    
			ans = step
			return
		}
		vis[cur] = true
		for i := 0; i < 8; i++ {
    
			for _, v := range dirs {
    
				temp := []byte(cur)
				if v != temp[i] {
    
					temp[i] = v
					cnt := string(temp)
					if ans != math.MaxInt32 || !mp[cnt] || vis[cnt] {
    
						continue
					}
					dfs(string(temp), step+1)
				}
			}
		}
		vis[cur] = false
	}
	dfs(start, 0)
	if ans == math.MaxInt32 {
    
		return -1
	}
	return ans
}