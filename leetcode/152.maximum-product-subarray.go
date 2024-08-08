package main

import (
	"fmt"
	"os"
)

// @leet start
func maxProduct(nums []int) int {
    dp_max := make([]int, len(nums))
    dp_min := make([]int, len(nums))
    dp_max[0], dp_min[0] = nums[0], nums[0]
    ans := nums[0]


    for i, _ := range dp_max {
        if i  == 0 {
            continue
        }

        maxprod := dp_max[i-1]
        minprod := dp_min[i-1]


        if nums[i] < 0 {
            maxprod, minprod = minprod, maxprod
        }

        dp_max[i] = max(nums[i], nums[i] * maxprod)
        dp_min[i] = min(nums[i], nums[i] * minprod) 
        ans = max(ans, dp_max[i])
    }

    return ans
}
// @leet end
