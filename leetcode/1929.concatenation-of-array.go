package main

// @leet start
func getConcatenation(nums []int) []int {
    results := make([]int, 2*len(nums), 2*len(nums))

    for i, _ := range results{
        results[i] = nums[i % len(nums)]
    }

    return results
}
// @leet end
