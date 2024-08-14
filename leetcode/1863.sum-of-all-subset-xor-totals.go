package main

// @leet start
func subsetXORSum(nums []int) int {
    sum := 0
    var backtrack func(idx int, xor int)
    backtrack = func(idx int, xor int){
        sum += xor

        for i:= idx; i< len(nums); i++{
            backtrack(i+1, xor ^ nums[i])
        }
    }

    backtrack(0, 0)

    return sum
}

// @leet end
