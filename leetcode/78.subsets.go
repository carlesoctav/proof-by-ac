package main

// @leet start
func subsets(nums []int) [][]int {
    ans := make([][]int, 0)
    curr := make([]int, 0)
    var backtracking func(idx int)

    backtracking = func(idx int){
        ans = append(ans, append([]int{}, curr...))
        if len(curr) == len(nums){
            return
        }
        for item_id := idx; item_id < len(nums); item_id++{
            curr = append(curr, nums[item_id])
            backtracking(item_id+1)
            curr = curr[: len(curr)-1]
        }
    }

    backtracking(0)
    return ans
}
// @leet end
