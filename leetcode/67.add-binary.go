package main

import (
	"fmt"
	"os"
	"strconv"
)

// @leet start
func addBinary(a string, b string) string {
    repeatZero := func (inp int) string {
        ans := ""
        for i:= 0; i < inp; i++ {
            ans += "0"
        }
        return ans 
    }
    maxbin := max(len(a), len(b))
    pad_a := repeatZero(maxbin - len(a)) + a
    pad_b := repeatZero(maxbin - len(b)) + b
    ans := "" 
    carry := 0


    for i := maxbin -1 ; i >= 0; i-- {
        a_int := int(pad_a[i]) - 48
        b_int := int(pad_b[i]) - 48
        if a_int + b_int + carry >= 2 {
            right := (a_int + b_int + carry) % 2
            ans = strconv.Itoa(right) + ans
            carry  = 1
        } else {
            right := (a_int + b_int + carry) % 2
            ans =strconv.Itoa((right)) + ans
            carry = 0
        }
    }

    if carry == 1{
        ans = strconv.Itoa(carry) + ans
    }
    return ans

}
// @leet end
