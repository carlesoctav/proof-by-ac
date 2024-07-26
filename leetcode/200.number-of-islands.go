package main

// @leet start
func dfs(i, j int, visited [][]bool, n, m int, grid [][]byte) {
    if visited[i][j] {
        return
    }

    visited[i][j] = true
    dr := []int{1, -1, 0, 0}
    dc := []int{0, 0, 1, -1}

    for mv := 0; mv < 4; mv++ {
        newI := i + dr[mv]
        newJ := j + dc[mv]
        if newI >= 0 && newI < n && newJ >= 0 && newJ < m && string(grid[newI][newJ]) == "1" {
            dfs(newI, newJ, visited, n, m, grid)
        }
    }
}

func numIslands(grid [][]byte) int {
    n := len(grid)
    m := len(grid[0])
    ans := 0
    visited := make([][]bool, n)

    for i := range visited {
        visited[i] = make([]bool, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if string(grid[i][j]) == "1" && !visited[i][j] {
                dfs(i, j, visited, n, m, grid)
                ans++
            }
        }
    }

    return ans
}
// @leet end
