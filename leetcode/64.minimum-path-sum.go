package main

// @leet start
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dr := []int{-1, 0}
    dc := []int{0, -1}
    var backtrack func(i, j int) int

    backtrack := func(i, j int)int {
        if i + dr[] 


    }


    return backtrack(m, n)
    
}
// @leet end
