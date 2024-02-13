package main

import (
	"math/rand"
)

// @leet start

func randomInt(start, end int) int {
    return start + rand.Intn(max-start)
}

func getParitionIndex(A []int, start, end int) int{
    pivot := A[randomInt(start, end+1)]
    pIndex := start
    for i := start; i<end; i++{
        if A[i]<=pivot {
            A[i], A[pIndex] = A[pIndex], A[i]
            pIndex++
        }

    }
    A[pIndex], A[end] = A[end], A[pIndex]
    return pIndex
}

func quickShort(nums []int, start int, end int){
    if start < end {
        partitionIndex := getParitionIndex(nums, start, end)

        quickShort(nums, start, partitionIndex-1)
        quickShort(nums, partitionIndex+1, end)
    }
}

func sortArray(nums []int) []int {
    quickShort(nums, 0, len(nums)-1)
    return nums 
}
// @leet end
