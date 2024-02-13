package main

// @leet start
func rob(nums []int) int {
    n := len(nums)+1
    dp := make([]int, n)

    dp[0] = nums[0]
    dp[1] = nums[1]

    for i:= 2 ; i < n-1 ; i++{
        dp[i] = max(dp[i-2]+nums[i], dp[i-1])
    }

    dp[n] = max(dp[n-1], dp[n-2])

    return dp[n]
}
// @leet end
