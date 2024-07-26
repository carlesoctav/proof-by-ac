package main

import (
	"fmt"
	"os"
)

// @leet start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2
    if len(A) > len(B){
        A,B = B,A
    }
    avg := (len(A)+len(B)+1)/2

    var Aleft, Aright float64
    var Bleft, Bright float64

    l, r := 0, len(A) -1
    for l<=r {
        i= l + (r-l)/2
        j = avg-i -2
        Aleft =
    }
}
// @leet end
