package main

// @leet start
func maxSubarraySumCircular(nums []int) int {
    dp := make([]int, 2 * len(nums))
    n := len(nums)
    ans := nums[0]
    dp[0] = nums[0]

    for i, _ := range dp {
        if i == 0 {
            continue
        }
        dp[i] = max(dp[i-1], nums[i % n] + dp[i-1])
        ans = max(dp[i], ans)
    }


    return ans
    
}
// @leet end
