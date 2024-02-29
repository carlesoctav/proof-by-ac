package main

// @leet start
func search(nums []int, target int) int {
    low := 0
    hi := len(nums) -1

    for low <= hi{
        mid := (low+ hi)/2
        if nums[mid] == target{
            return mid
        } else if nums[mid] < target{
            low = mid +1
        } else {
            hi = mid -1
        }
    }

    return -1

    
}
// @leet end
