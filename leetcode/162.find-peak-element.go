package main

// @leet start
func findPeakElement(nums []int) int {
    condition := func(mid int ) bool {
        if mid == 0{
            return true
        }

        if nums[mid] > nums[mid -1]{
            return true
        }

        return false
    }
    binSearch := func(nums []int) int{
        ans := -1
        l, r := 0, len(nums) -1
        for l<=r{
            mid := l + (r-l)/2
            if condition(mid){
                ans = mid
                l = mid +1
            } else {
                r = mid - 1
            }
        }
        return ans
    }

    return binSearch(nums)
}
// @leet end
