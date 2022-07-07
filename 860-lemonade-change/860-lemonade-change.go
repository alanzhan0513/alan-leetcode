func lemonadeChange(bills []int) bool {
    coins := make(map[int]int)
    coins[5] = 0
    coins[10] = 0
    coins[20] = 0
    for _, bill := range bills {
        coins[bill]++

        if !canExchange(bill - 5, coins) {
            return false
        }
    }
    return true
}

func canExchange(amount int, coins map[int]int) bool {
    for _, x := range []int{20, 10, 5} {
        for amount >= x && coins[x] > 0 {
            amount -= x
            coins[x]--
        }
    }
    return amount == 0
}