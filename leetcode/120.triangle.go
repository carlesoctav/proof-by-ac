package main

import (
	"fmt"
	"os"
)

// @leet start
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
            if j == 0 {
                dp[i][j] = (dp[i-1][j] + triangle[i][j])
            } else if j == len(triangle[i]) -1 {
                dp[i][j] = (dp[i-1][j-1]+ triangle[i][j])

            } else {
                dp[i][j] = min(dp[i-1][j] + triangle[i][j], dp[i-1][j-1]+ triangle[i][j]) 

            }
		}
	}

	min_val := int(1e9)

	for _, val := range dp[len(triangle)-1] {
		min_val = min(min_val, val)

	}
	return min_val
}

// @leet end

