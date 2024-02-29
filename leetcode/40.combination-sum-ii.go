package main

import "sort"

// @leet start
func combinationSum2(candidates []int, target int) [][]int {

    sort.Ints(candidates)
    ans := [][]int{}
    curr := []int{}

    var backtrack func(idx int)

    sum := func(curr []int) int{
        ans := 0
        for _, value := range curr{
            ans+=value
        }

        return ans
    }

    backtrack = func(idx int){
        if sum(curr) == target{
            ans = append(ans, append([]int{}, curr...))
            return
        } 

        if sum(curr) > target{
            return 
        }
        
        for i:=idx; i<len(candidates); i++{
            if i > idx && candidates[i]  == candidates[i-1]{
                continue
            }
            curr = append(curr, candidates[i])
            backtrack(i+1)
            curr = curr[:len(curr)-1]
        }

    }
    backtrack(0)
    return ans
}
// @leet end
