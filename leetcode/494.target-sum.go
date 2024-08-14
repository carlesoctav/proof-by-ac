package main

// @leet start
func findTargetSumWays(nums []int, target int) int {
    count := 0
    var backtrack func(idx int, sum int)
    backtrack = func(idx int, sum int){
        if idx == len(nums){
            if sum == target {
                count++
            }
            return 
        }
            backtrack(idx+1, sum + nums[idx])
            backtrack(idx+1,  sum - nums[idx])
    }

    backtrack(0, 0)
    return count
}
// @leet end
