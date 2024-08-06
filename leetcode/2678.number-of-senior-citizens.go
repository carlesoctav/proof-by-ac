package main

import "strconv"

// @leet start
func countSeniors(details []string) int {

    count := 0
    for _, val := range details {
        if  v, _ :=strconv.Atoi(val[11:13]); v > 60{
            count++
        }
    }

    return count
    
}
// @leet end
