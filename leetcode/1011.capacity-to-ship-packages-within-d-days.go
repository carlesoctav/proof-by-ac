package main

import (
	"fmt"
	"os"
)

// @leet start
func shipWithinDays(weights []int, days int) int {

    sum := func( w []int) int {
        ans := 0
        for _, value := range w{
            ans += value
        }
        return ans
    }

    left, right := 0, sum(weights)
    condition := func(mid int) bool{
        d := 1
        currContainer :=0
        for _, value := range weights {

            if mid < value {
                return  false
            }

            if currContainer + value <= mid {
                currContainer += value
            } else{
                d++
                currContainer = value
            }
        }

        return d <=  days
    }

    for left < right {
        mid := left + (right - left) /2
        if condition(mid){
            right = mid
            
        } else {
           left = mid +1 
        }
    }
    return left
}
// @leet end
