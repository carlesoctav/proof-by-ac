package main

import "fmt"

// @leet start
func lengthOfLongestSubstring(s string) int {
    max := func(a, b int) int{
        if a > b {
            return a
        } else {
            return b
        }
    }
    charSet := make(map[byte]bool) 
    l := 0
    ans := 0
    for r := range s{
        for charSet[s[r]] {
            delete(charSet, s[l])
            l++
        }
        charSet[s[r]] = true
        ans = max(ans, r-l+1)
    }

    return ans
}
// @leet end
