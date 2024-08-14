package main

// @leet start
func combine(n int, k int) [][]int {
    var backtrack func(idx int, curr []int)
    ans := make([][]int, 0)

    backtrack = func(idx int, curr []int) {
        if len(curr) == k {
            ans = append(ans, append([]int{}, curr...))
            return 
        }

        for i := idx; i < n; i++{
            backtrack(i+1, append(curr, i+1))
        }
    }

    backtrack(0, []int{})
    return ans
}
// @leet end
