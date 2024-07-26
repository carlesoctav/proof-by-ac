package main

// @leet start
func missingNumber(nums []int) int {
    res := 0
    for i := 1; i <= len(nums); i++{
        res += (i - nums[i-1])
    }

    return res
    
}
// @leet end
