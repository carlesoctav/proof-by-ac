package main

// @leet start
func minSwaps(nums []int) int {
    total_one := 0
    max_window := 0
    window := 0
    for _, val := range nums {
        if val == 1{
            total_one +=1
        }
    }

    i := 0 
    for j := 0; j < 2 * len(nums); j++ {
        if nums[j % len(nums)] == 1{
            window+=1
        }
        if j -i + 1 > total_one {
            window -= nums[i % len(nums)] 
            i++
        }
        max_window = max(window, max_window)
    }


    return total_one - max_window
}
// @leet end
