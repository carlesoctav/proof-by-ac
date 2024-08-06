package main

import (
	"fmt"
	"os"
	"sort"
)

// @leet start
func minimumPushes(word string) int {
    occurance := make([]int, 26)
    for _, val:= range word {
        occurance[val - 'a'] +=1
    }

    sort.Slice(occurance, func (i, j int) bool {
        return occurance[i] > occurance[j]
    })
    fmt.Fprintf(os.Stdout, "DEBUGPRINT[1]: 3016.minimum-number-of-pushes-to-type-word-ii.go:9: occurance=%+v\n", occurance)

    ans := 0

    for i, val := range occurance {
        mult := (i / 8) + 1
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[2]: 3016.minimum-number-of-pushes-to-type-word-ii.go:24: mult=%+v\n", mult)
        ans += (val * mult)
    }


    return ans
    
}
// @leet end
