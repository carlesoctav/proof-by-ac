package main

import (
	"fmt"
	"os"
	"strconv"
)

// @leet start
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		for j := i -1 ; i-j <= 2 && j >= 0; j-- {
            if s[j] == '0'{
                continue
            }
			conv, _ := strconv.Atoi(s[j:i])
			fmt.Fprintf(os.Stdout, "DEBUGPRINT[3]: 91.decode-ways.go:14: conv=%+v\n", conv)
			if 1 <= conv && conv <= 26 {
				dp[i] += dp[j]
			}
		}
	}
	return dp[len(s)]

}

// @leet end
