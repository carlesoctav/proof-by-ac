package main

import (
	"fmt"
	"os"
)

// @leet start
func dfs_max_area(i, j int, visited [][]bool, grid [][]int, count *int, m, n int){
    
    if visited[i][j]{
        return
    }

    visited[i][j] = true
    *count +=1

    dr := []int{0,0,1,-1}
    dc := []int {1,-1,0,0}

    for mv := 0; mv < 4; mv++{
        newI := i + dr[mv]
        newJ := j + dc[mv]
        if newI >=0 && newI < m && newJ >=0 && newJ < n && grid[newI][newJ] == 1 {
            dfs_max_area(newI,newJ, visited, grid, count, m, n)
        }

    }
}
func maxAreaOfIsland(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    visited := make([][]bool, m)
    for i := range visited{
        visited[i] = make([]bool, n)
    }

    ans := 0
    count := 0
    ifVisited := 0

    for i := 0; i < m; i++{
        for j :=0; j < n; j++{
            if !visited[i][j] && grid[i][j] == 1{
                ifVisited++
                dfs_max_area(i,j, visited, grid, &count, m, n)
                ans = max(ans, count)
                count = 0
            }
        }
    }
    fmt.Fprintf(os.Stdout, "DEBUGPRINT[2]: 695.max-area-of-island.go:39: ifVisited=%+v\n", ifVisited)

    return ans
    
}
// @leet end
