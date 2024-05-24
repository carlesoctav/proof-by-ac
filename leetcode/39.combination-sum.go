package main

// @leet start
func combinationSum(candidates []int, target int) [][]int {
    ans := make([][]int,0)
    curr := make([]int, 0)
    sum := func(arr []int)int {
        output := 0
        for _, v := range arr{
            output+=v
        }

        return output
    }
    var backtracking func(idx int)
    backtracking = func(idx int){
        if sum(curr) == target{
            ans = append(ans, append([]int{}, curr...))
        }

        if sum(curr) > target{
            return
        }


        for curr_id := idx; curr_id< len(candidates); curr_id++{
            curr = append(curr, candidates[curr_id])
            backtracking(curr_id)
            curr = curr[:len(curr)-1]
        }
    }

    backtracking(0)
    return ans
}
// @leet end
