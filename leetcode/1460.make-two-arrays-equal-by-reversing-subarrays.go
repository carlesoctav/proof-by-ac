package main

// @leet start
func canBeEqual(target []int, arr []int) bool {
    mapping := make(map[int]int)
    for _, val := range target {
        if _, ok := mapping[val]; !ok {
            mapping[val] = 1 
        } else {
            mapping[val]+=1
        }
    }


    for _, val := range  arr {
        if _, ok := mapping[val]; ok {
            mapping[val] -= 1

        } else {
            return false
        }
    }
    for _, val := range mapping {
        if  val  != 0 {
            return false
        }
    }
    return true
}
// @leet end
