package main

// @leet start
func singleNumber(nums []int) int {

    res := 0
    for _, value := range nums{
        res ^= value
    }

    return res
    
}
// @leet end
