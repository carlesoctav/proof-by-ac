package main

import "sort"

// @leet start

type Pair struct{
    val int
    idx int
}

func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)

    ans := [][]int{}
    for nums1Idx := 0; nums1Idx < n-2; nums1Idx++{
        if nums1Idx > 0 && nums[nums1Idx] == nums[nums1Idx-1]{
            continue
        }

        nums2Idx := nums1Idx+1
        nums3Idx := n-1

        for nums2Idx < nums3Idx {
            sum:= nums[nums1Idx] + nums[nums2Idx] + nums[nums3Idx]

            if sum < 0{
                nums2Idx++
            } else if sum >0 {
               nums3Idx-- 
            } else {
                ans = append(ans, []int{nums[nums1Idx], nums[nums2Idx], nums[nums3Idx]})
                nums2Idx++

                for nums2Idx < nums3Idx && nums[nums2Idx] == nums[nums2Idx-1]{
                    nums2Idx++
                }

            }
        }
    }
    return ans
}
