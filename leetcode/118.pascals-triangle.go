package main

import (
	"fmt"
	"os"
)

// @leet start
func generate(numRows int) [][]int {

    if numRows == 1 {
            return [][]int{{1}}
        }

    ans := [][]int{{1}}
    for i := 2; i <= numRows; i++{
        ansI := []int{}
        ansI = append(ansI, 1)
        prevPascal := len(ans[len(ans) -1 ])
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[1]: 118.pascals-triangle.go:13: prevPascal=%+v\n", prevPascal)
        for j := 0; j< prevPascal -1 ; j++ {
            ansI = append(ansI, ans[prevPascal][j]+ ans[prevPascal][j+1])
        }
        ansI = append(ansI, 1)
        ans = append(ans, ansI)
    }
    return ans
}
// @leet end
