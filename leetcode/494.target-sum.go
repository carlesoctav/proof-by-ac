package main

// @leet start
func findTargetSumWays(nums []int, target int) int {
    dp := make([][]int, len(nums)+1)
    for i := range dp {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1

    for i := 1; i <= len(nums); i++ {
        for j := 1; j <= target; j++{
            if j - nums[i-1] >= 0{
                dp[i][j] += dp[i-1][j - nums[i-1]]
            }
            if j + nums[i-1] <= target {
                dp[i][j] += dp[i-1][j + nums[i-1]]
            }
            
        }
    }

    return dp[len(nums)][target]
    
}
// @leet end
