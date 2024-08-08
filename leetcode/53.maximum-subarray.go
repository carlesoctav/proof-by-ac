package main

// @leet start
func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    ans := nums[0]
    dp[0] = nums[0]


    for i:= 1; i< len(nums); i++{
        dp[i] = max(nums[i], dp[i-1]+nums[i])
        ans = max(dp[i], ans)
    }

    return ans
}
// @leet end
