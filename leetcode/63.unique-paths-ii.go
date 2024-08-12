package main

import (
	"fmt"
	"os"
)

// @leet start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))


	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	dp[0][0] = 1
	dr := []int{-1, 0}
	dc := []int{0, -1}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			for k := 0; k < 2; k++ {
                if obstacleGrid[i][j] == 1{
                    dp[i][j] = 0
                    continue
                }
				if i+dr[k] < 0 || i+dr[k] >= len(obstacleGrid) || j+dc[k] < 0 || j+dc[k] >= len(obstacleGrid[i]) {
                    continue
				}


                dp[i][j]+=dp[i+dr[k]][j+dc[k]]
			}
		}
	}

    fmt.Fprintf(os.Stdout, "DEBUGPRINT[1]: 63.unique-paths-ii.go:32: dp=%+v\n", dp)
    return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

// @leet end

