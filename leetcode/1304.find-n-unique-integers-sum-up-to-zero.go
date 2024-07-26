package main

// @leet start
func sumZero(n int) []int {
    ans := make([]int, 0)
    if n % 2 == 1 {
        ans = append(ans, 0) 
        n = n -1
    }

    for i:= 0; i<n/2;i++{
        ans = append(ans, i+1)
        ans = append(ans, -( i+1 ))
    }

    return ans
}
// @leet end
