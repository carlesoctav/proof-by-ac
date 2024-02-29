package main

import (
	"fmt"
	"os"
	"sort"
)

// @leet start
func subsetsWithDup(nums []int) [][]int {
    ans := [][]int{}
    curr := []int{}
    sort.Ints(nums)
    var backtrack func (idx int) 

    backtrack = func(idx int){
        ans = append(ans, append([]int{}, curr...))

        if idx == len(nums){
            return
        }

        for i:= idx; i<len(nums); i++{
            if i > idx && nums[i] == nums[i-1]{
                continue;
            }

            curr = append(curr, nums[i])
            backtrack(i+1)
            curr = curr[:len(curr)-1]
            
        }
    }
    backtrack(0)
    return ans
}
// @leet end
