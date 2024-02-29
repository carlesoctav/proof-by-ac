package main

// @leet start
func findKthNumber(m int, n int, k int) int {


    lo, hi := 1, m * n

    condition := func(mid int) bool{
        amount :=0
        for i:= 1; i<=m; i++{
            amount += min(n, mid /i)
        }
        return amount >= k
    }

    for lo < hi {
        mid := lo + (hi - lo)/2
        if condition(mid){
           hi = mid 
        } else {
            lo = mid +1
        }
    }

    return lo
}
// @leet end
