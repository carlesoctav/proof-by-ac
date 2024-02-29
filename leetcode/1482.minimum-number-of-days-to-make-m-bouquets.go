package main

import (
	"fmt"
	"os"
	"slices"
)

// @leet start
func minDays(bloomDay []int, m int, k int) int {

    condition := func(mid int) bool {
        count := 0
        bk := 0
        for _, value := range bloomDay{
            if mid - value >=0  {
                count+=1
            } else {
                count = 0
            }

            if count == k {
                bk++
                count = 0
            }

        }
        fmt.Println(bk)
        return bk >= m
    }

    lo , hi := slices.Min(bloomDay), slices.Max(bloomDay)+1

    for lo < hi {
        mid := lo + (hi - lo)/2
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[2]: 1482.minimum-number-of-days-to-make-m-bouquets.go:32: mid=%+v\n", mid)
        if (condition(mid)){
            hi = mid
        } else {
            lo = mid +1
        }
    }


    if lo != slices.Max(bloomDay)+1{
        return lo
    } else {
        return -1
    }
    
}
// @leet end
