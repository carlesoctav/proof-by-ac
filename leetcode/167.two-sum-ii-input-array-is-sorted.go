package main

// @leet start
func twoSumNew(numbers []int, target int) []int {
    left_pointer := 0
    right_pointer := len(numbers)-1

    for {
        if numbers[left_pointer]+ numbers[right_pointer] < target{
            left_pointer++
        } else if numbers[left_pointer] + numbers[right_pointer]>target {
            right_pointer--
        } else {
            return []int{left_pointer+1, right_pointer+1}
        }

    }

    
}
// @leet end
