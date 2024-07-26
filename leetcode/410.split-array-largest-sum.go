package main

import "slices"

// @leet start
func splitArray(nums []int, k int) int {
    sum := func (nums []int) int {
        ans := 0
        for _, value := range nums{
            ans += value
        }

        return ans
    }


    condition := func(minMax int) bool {
        partition := 1
        currSum := 0
        for _, value := range nums{
            if currSum + value <= minMax {
                currSum += value
            } else {
                partition++
                currSum = value
            }
        }

        return  partition <= k
    }
    lo, hi := slices.Max(nums), sum(nums)

    for lo < hi {
        mid := lo + (hi -lo)/2
        if condition(mid){
            hi = mid
        } else {
            lo = mid +1
        }
    }
    return lo
}
// @leet end
