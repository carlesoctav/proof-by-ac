package main

// @leet start
func searchInsert(nums []int, target int) int {

    l, r := 0, len(nums)

    condition := func(mid int) bool {
        return nums[mid] >= target
    }

    for l< r{
        mid := l + (r-l)/2

        if condition(mid){
            r = mid
        } else {
            l = mid +1
        }
    }
    return l
}
// @leet end
