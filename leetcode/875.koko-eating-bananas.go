package main

import "slices"

// @leet start
func minEatingSpeed(piles []int, h int) int {

    condition := func(mid int) bool {
        total_hour := 0

        for _, value := range piles {
            total_hour += value / mid
            if value % mid != 0 {
                total_hour +=1
            }
        }
        return total_hour <= h
    }

    lo, hi := 1, slices.Max(piles)

    for lo < hi{
        m := lo + (hi - lo) / 2
        if condition(m){
            hi = m
        } else {
            lo = m +1
        }
    }
    return lo
}
// @leet end
