package main

import "strings"

// @leet start
func wordPattern(pattern string, s string) bool {
    stringList := strings.Split(s, " ")
    mapPatternToString := make(map[string]string)
    mapStringToPattern := make(map[string]string)

    if len(stringList) != len(pattern){
        return false
    }

    for i, val := range pattern{

        if whatVal, ok := mapStringToPattern[stringList[i]]; ok {
            if whatVal != string(val) {
                return false
            }
        }
        
        if val2, ok := mapPatternToString[string(val)]; !ok{
            mapPatternToString[string(val)] = stringList[i]
            mapStringToPattern[stringList[i]] = string(val)

        } else {
            if val2 != stringList[i]{
                return false
            }
        }
    }


    return true
    
}
// @leet end
