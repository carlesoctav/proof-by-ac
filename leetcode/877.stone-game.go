package main

// @leet start
func stoneGame(piles []int) bool { 
    var backtrack func(i, j int)int
    dp := make([][]int, len(piles))
    for i := range dp {
        dp[i] = make([]int, len(piles))
    }
    backtrack = func(i, j int)int{
        if i > j{
            return 0 
        }

        if dp[i][j] != 0{
            return dp[i][j]
        }

        length := j - i
        if length % 2 == 0 {
            takeLeft := piles[i]
            takeRight := piles[j]
            dp[i][j] = max(backtrack(i+1, j)+ takeLeft, backtrack(i, j-1) + takeRight)
        } else {
            takeLeft := 0 
            takeRight := 0 
            dp[i][j] = max(backtrack(i+1, j)+ takeLeft, backtrack(i, j-1) + takeRight)
        }

        return dp[i][j]
    }

    sum := func(curr []int) int {
        a := 0
        for _, val := range curr {
            a+=val
        }
        return a
    }

    backtrack(0, len(piles)-1)
    return dp[0][len(piles) - 1] > sum(piles) /2
    
}
// @leet end
