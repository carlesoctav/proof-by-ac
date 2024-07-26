package main

import "fmt"

// @leet start
func myPow(x float64, n int) float64 {
    AbsSign := func(n int) (abs int, sign bool){
        if n < 0{
            abs = -n
            sign = false
            return
        }

        abs = n
        sign = true
        return 
    }

    absN, signN := AbsSign(n)

    ans := 1.0
    incrementPow := x
    for  absN > 0{
        fmt.Println(absN)
        if absN % 2 == 1 {
            ans *= incrementPow
            incrementPow *= incrementPow

        } else {
            incrementPow *= incrementPow
        }

        absN >>=1
    }

    if signN {
        return ans
    } else {
        return 1/ans
    }
}
// @leet end
