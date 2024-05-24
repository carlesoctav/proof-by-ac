package main

// @leet start
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    min := func(x, y int) int {
        if x <y{
            return x
        } else {
            return y
        }
    }
    for i, _:= range dp{
        dp[i] = amount + 1
    }
    dp[0] = 0
    for i:= 1;i < len(dp); i++{
        for _, c:= range coins{
            if c <= i{
                dp[i] = min(dp[i] , dp[i-c] + 1)
            }
        }
    }

    if dp[amount] != amount + 1{
        return dp[amount]
    } else{
        return -1
    }

}
// @leet end
