package main

// @leet start
func maxProfit(prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }

    dp := make([][3]int, n)

    dp[0][0] = 0
    dp[0][1] = -prices[0]
    dp[0][2] = 0

    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

        dp[i][1] = max(dp[i-1][1], dp[i-1][2]-prices[i])

        dp[i][2] = dp[i-1][0]
    }

    return max(dp[n-1][0], dp[n-1][2])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// @leet end
