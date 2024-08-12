package main

// @leet start
func canPartition(nums []int) bool {
    sum := 0
    for _, val := range nums {
        sum += val
    }
    if sum % 2 == 1 {
        return false
    }
    k := sum / 2

    dp := make([][]bool, len(nums) + 1)

    for i := range dp {
        dp[i] = make([]bool, k+1)
        dp[i][0] = true
    }

    for i := 1; i< len(nums)+1 ; i++ {
        for j:= 1; j < k+1; j++ {
            dp[i][j] = dp[i-1][j]
            if j - nums[i-1] >= 0 {
                dp[i][j] = dp[i-1][ j - nums[i-1]] || dp[i][j]
            }

        }
    }

    return dp[len(nums)][k]
}
// @leet end
