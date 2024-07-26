package main

// @leet start
func permute(nums []int) [][]int {

    ans := make([][]int, 0)
    curr := make([]int, 0)
    var backtrack func (id int)
    chosen := make(map[int]bool)
    backtrack = func(i int){
        if len(curr) == len(nums){
            ans = append(ans, append([]int{}, curr...))
        }

        for i := 0; i < len(nums); i++{
            if chosen[i] == false {
                chosen[i] = true
                curr = append(curr, nums[i])
                backtrack(i+1)
                curr = curr[:len(curr)-1]
                chosen[i] = false
            }
        }
    }

    backtrack(0)
    return ans

}
// @leet end
