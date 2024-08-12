package main

import (
	"fmt"
	"os"
	"strings"
)

// @leet start
func wordBreak(s string, wordDict []string) []string {
    var backtrack func(idx int)
    word := make(map[string]bool)
    ans := make([]string,0)
    curr := make([]string, 0)

    for _, val := range wordDict {
        word[val] = true
    }

    backtrack = func(idx int) {
        if idx == len(s) {
            ans = append(ans, strings.Join(curr, " "))
            return
        }

        for i:= idx; i < len(s); i++ {
            _, ok := word[s[idx:i+1]]
            if ok {
                curr = append(curr, s[idx:i+1])
                backtrack(i+1)
                curr = curr[0: len(curr) - 1]
            }
        }
    }

    backtrack(0)

    return ans

}
// @leet end
