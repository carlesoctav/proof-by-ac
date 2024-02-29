package main

// @leet start
func permute(nums []int) [][]int {

    visited := make([]bool, len(nums))
    curr := []int{}
    ans := [][]int{}

    var backtrack func()

    backtrack = func(){

        if len(curr) == len(nums){
            ans = append(ans, append([]int{}, curr...))
            return 
        }

        for i:= 0; i< len(nums); i++{

            if !visited[i]{
                visited[i] = true
                curr = append(curr, nums[i])
                backtrack()
                visited[i] = false
                curr = curr[: len(curr)-1]
            }

        }

    }
    backtrack()

    return ans
}
// @leet end
