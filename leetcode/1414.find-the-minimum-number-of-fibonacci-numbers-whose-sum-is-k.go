package main

import (
	"fmt"
	"os"
)

// @leet start
func findMinFibonacciNumbers(k int) int {
    fib := make([]int, 0)
    prev := 1
    curr := 1
    condition := func(mid, target int) bool {
        return fib[mid] <= target
    }
    search := func(i,j, target int) int {
        ans := 0
        for i<= j {
            mid := i + (j-i)/2
            if condition(mid, target){
                ans = mid
                i = mid + 1
            } else {
                j = mid -1 
            }
        }
        return ans
    }

    for {
        next := prev + curr
        fib = append(fib, next)
        prev = curr
        curr = next

        if curr >= 1e9 {
            break
        }
    }

    count := 0
    for k > 0 {
        cool := fib[search(0, len(fib)-1, k)]
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[2]: 1414.find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k.go:42: cool=%+v\n", cool)
        k -= fib[search(0, len(fib) - 1, k)]
        count++
    }

    return count
}
// @leet end
