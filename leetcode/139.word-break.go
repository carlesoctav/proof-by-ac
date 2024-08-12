package main

// @leet start
func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    word := make(map[string]bool)

    for _, val := range wordDict{
        word[val] = true
    }
    dp[0] = true 

    for i := 1; i <=len(s); i++{
        for j := i ; j >= 0; j--{
            _, ok := word[s[j:i]]
            dp[i] = (dp[j] && ok) || dp[i] 
        }
    }

    return dp[len(s)]
    
}
// @leet end
