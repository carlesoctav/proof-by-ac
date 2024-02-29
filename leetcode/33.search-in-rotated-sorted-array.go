package main

import (
	"fmt"
	"os"
)

// @leet start
func search(nums []int, target int) int {

    binSearch := func(lo, hi, target int) int {
        for lo <= hi {
            mid := lo + (hi - lo)/2
            fmt.Fprintf(os.Stdout, "DEBUGPRINT[9]: 33.search-in-rotated-sorted-array.go:14: nums[mid]=%+v\n", nums[mid])
            if nums[mid] == target {
                fmt.Fprintf(os.Stdout, "DEBUGPRINT[8]: 33.search-in-rotated-sorted-array.go:16: mid=%+v\n", mid)
                return mid
            } else if nums[mid] < target{
                lo = mid +1
            } else {
                hi = mid -1
            }
        }

        return -1
    }


    lo, hi := 0, len(nums) -1

    for lo < hi {
        mid := lo + ( hi - lo)/2

        if nums[mid] < nums[0]{
            hi = mid 
        } else {
            lo = mid + 1 
        }
    }

    fmt.Fprintf(os.Stdout, "DEBUGPRINT[7]: 33.search-in-rotated-sorted-array.go:39: lo=%+v\n", lo)
    ans := max(-1, binSearch(0, lo, target), binSearch(lo, len(nums) - 1, target))
    return ans
}
// @leet end
