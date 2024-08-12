package main

import "sort"

// @leet start
func permuteUnique(nums []int) [][]int {
    var backtrack func(curr []int)
    sort.Ints(nums)
    ans := make([][]int, 0)
    choose := make(map[int]int)
    for _, val := range nums {
        _, ok := choose[val]
        if !ok {
            choose[val] = 1
        } else {
            choose[val]++
        }
    }
    backtrack = func(curr []int){
        if len(curr) == len(nums){
            ans = append(ans, append([]int{}, curr...))

        }

        for i := 0; i < len(nums); i++{

            if i != 0 && nums[i] == nums[i-1] && choose[nums[i]] != 0 {
                continue
            }

            if choose[nums[i]] != 0{
                choose[nums[i]]--
                backtrack(append(curr, nums[i]))
                choose[nums[i]]++ 
            }
        }
    }

    backtrack([]int{})
    return ans
}
// @leet end
