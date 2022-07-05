func coinChange(coins []int, amount int) int {
    var dp =make([]int,amount+1)
    for i:=0;i<amount+1;i++{
        dp[i]=amount+1
    }
    dp[0]=0
    for i:=1;i<=amount;i++{
        for j:=0;j<len(coins);j++{
            if coins[j]==i{
                dp[i]=1
            }else if i>coins[j]{
                if dp[i]>dp[i-coins[j]]+1{
                    dp[i]=dp[i-coins[j]]+1
                }
            }
        }
    }
    if dp[amount]>amount{
        return -1
    }else{
        return dp[amount]
    }
}