package main

import "fmt"

// @leet start
func productExceptSelf(nums []int) []int {
    prefix_prod := make([]int, len(nums), len(nums))
    suffix_prod := make([]int, len(nums), len(nums))
 
    for i, _ := range nums {
        if i ==0 {
            prefix_prod[i] = 1
            continue
        }
        
        prefix_prod[i] = prefix_prod[i-1] * nums[i-1]
    }
    fmt.Printf("%v", prefix_prod)
    for i:= len(nums)-1 ; i>=0; i-- {

        if i == len(nums)-1{
            suffix_prod[i] = 1
            continue
        }
        suffix_prod[i] = suffix_prod[i+1]* nums[i+1]
    }
    fmt.Printf("%v", suffix_prod)
    
    sol := []int{}

    for i:=0 ; i< len(nums); i++{
        sol = append(sol, prefix_prod[i]* suffix_prod[i])
    }

     return sol
}
// @leet end
