package main

import (
	"fmt"
	"os"
)

// @leet start
func findMin(nums []int) int {


    lo, hi := 0, len(nums) -1

    for lo < hi {
        mid := lo + ( hi - lo)/2
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[5]: 153.find-minimum-in-rotated-sorted-array.go:16: nums[mid]=%+v\n", nums[mid])
        if nums[mid] < nums[0]{
            hi = mid
        } else {
            lo = mid +1
        }
    }
    return min(nums[lo], nums[0], nums[len(nums) -1 ])
    
}
// @leet end
