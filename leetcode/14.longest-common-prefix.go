package main

import "fmt"

// @leet start
func longestCommonPrefix(strs []string) string {
    getShortestString := func(strs []string) int {
        value := 999
        idx := 0
        for i, str := range strs {
            fmt.Println(str, len(str))
            if value > len(str){
                value = len(str)
                idx = i
            }
        }
        return idx
    }
    shortIndex := getShortestString(strs)
    fmt.Println(shortIndex)
    ans := []rune{}
    for i:= 0; i<len(strs[shortIndex]);i++{
        for _, str := range strs{
            if str[i] != strs[shortIndex][i]{
                return string(ans)
            }
        }

        ans = append(ans, rune(strs[shortIndex][i]))
    }
    return string(ans)
}
// @leet end
