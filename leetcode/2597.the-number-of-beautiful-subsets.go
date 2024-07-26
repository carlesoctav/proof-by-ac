package main

import (
	"fmt"
	"os"
)

// @leet start
func beautifulSubsets(nums []int, k int) int {
    ans := 0
    subset := make([][]int , 0)
    curr := make([]int, 0)
    var backtracking func(id int)
    condition := func(curr []int, anchor int) bool {
        for _, v := range curr{
            if anchor - v == k ||  v - anchor == k{
                return false 
            }
        }

        return true

    }
    backtracking = func(id int){


        for i := id; i< len(nums); i++{
            if condition(curr, nums[i]){
                ans+=1
                curr = append(curr, nums[i])
                subset = append(subset, append([]int{}, curr...))
                backtracking(i+1)
                curr = curr[0: len(curr) -1] 
            }
        }
    }

    backtracking(0)
    return ans
}
// @leet end
