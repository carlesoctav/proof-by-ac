package main

// @leet start
func uniquePaths(m int, n int) int {

    var backtrack func(m, n int) int

    backtrack = func(m, n int) int {
        if m == 0 || n == 0{
            return 1
        }
        return backtrack(m, n-1) + backtrack(m-1, n)
    }

    dp := make([][]int,m)

    for i := range dp{
        dp[i] = make([]int, n)
    }
    
    dp[0][0] = 1

    for i:= 0; i< m; i++ {
        for j:= 0; j<n; j++{
            if j-1 >=0 {
                dp[i][j] += dp[i][j-1]
            }

            if i - 1 >= 0 {
                dp[i][j] +=dp[i-1][j]
            }
        }
    }

    return dp[m-1][n-1]
    
}
// @leet end
