package main

import (
	"math"
	"strconv"
)

// @leet start
func hammingWeight(num uint32) int {

    ans := num & math.MaxUint32
    binary_rep := strconv.FormatUint(uint64(ans), 2)
    count := 0

    for _, val := range binary_rep{
        if val == '1'{
            count +=1
        }

    }

    return count

    
    
}
// @leet end
