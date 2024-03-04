package main

import (
	"fmt"
	"os"
)

// @leet start
func searchMatrix(matrix [][]int, target int) bool {
    lo, hi := 0, len(matrix) -1
    searchIdx := -1
    for lo <= hi {
        mid := lo + ( hi - lo )/2
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[14]: 74.search-a-2d-matrix.go:13: mid=%+v\n", mid)
        fmt.Fprintf(os.Stdout, "DEBUGPRINT[16]: 74.search-a-2d-matrix.go:16: matrix[mid][0]=%+v\n", matrix[mid][0])
        if target >= matrix[mid][0] {
            searchIdx = mid
            lo = mid +1
        } else {
            hi = mid -1
        }
    }
    fmt.Fprintf(os.Stdout, "DEBUGPRINT[13]: 74.search-a-2d-matrix.go:24: searchIdx=%+v\n", searchIdx)

    if searchIdx == -1 {
       return false 
    }

    lo, hi = 0, len(matrix[0]) -1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if target == matrix[searchIdx][mid] {
            return true
        } else if target < matrix[searchIdx][mid]{
            hi = mid -1
        } else {
            lo = mid +1
        }
    }

    return false
    
}
// @leet end
