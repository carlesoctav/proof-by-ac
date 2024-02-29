package main

// @leet start
func combinationSum(candidates []int, target int) [][]int {
    var backtrack func(idx int)
    ans := [][]int{}
    curr := []int{}

    sum := func(curr []int) int {
        ans := 0
        for _, val := range curr {
            ans+= val
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

        for i := idx; i<len(candidates); i++{
            curr = append(curr, candidates[i])
            backtrack(i)
            curr = curr[:len(curr)-1]
        }
        
    }

    backtrack(0)
    
    return ans
}
// @leet end
