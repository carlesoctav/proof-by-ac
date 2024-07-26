package main

import "fmt"

// @leet start
func replaceElements(arr []int) []int {
    greatest := make([]int, len(arr), len(arr))

    if len(arr) ==1{
        return []int{-1}
    }

    if len(arr) == 2{
        return []int{arr[1], -1}
    }


    greatest[len(arr)-1] = -1
    greatest[len(arr)-2] = arr[len(arr)-1]
    for i:= len(arr)-3; i>=0;i--{
        greatest[i] = max(greatest[i+1], arr[i+1])
    }


    return greatest

}
// @leet end
