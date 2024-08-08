package main

// @leet start
func rob(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    if len(nums) == 1 {
        return nums[0]
    }
    dp[1] = nums[1]
    ans := max(dp[0], dp[1]) 


    for i, _ := range dp {
        if i == 0 || i == 1 {
            continue
        }

        for k:=2 ; k <= i; k++{
            dp[i] = max(nums[i], nums[i]+ dp[i-k], dp[i])
        }
        ans = max(ans, dp[i])
    } 

    return ans 
}
// @leet end
