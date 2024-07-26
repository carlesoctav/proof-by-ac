package main

import (
	"fmt"
	"os"
)

// @leet start
func numOfSubarrays(arr []int, k int, threshold int) int {
	sum := func(arr []int) int {
		total := 0
		for _, v := range arr {
			total += v
		}
		return total
	}

    count := 0
    running_sum := sum(arr[:k+1])
    i, j := 0, k

    for j < len(arr){
        if float64(running_sum) / float64(k) >= float64(threshold){
            count++
        }
        running_sum -= arr[i]
        i++
        j++
        if j < len(arr){
            running_sum += arr[j]
        }
    }

    return count
}
// @leet end
