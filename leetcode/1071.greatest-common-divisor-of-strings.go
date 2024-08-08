package main

import "os"

// @leet start
func gcdOfStrings(str1 string, str2 string) string {
    var commonstr string
    var longstr string
    if len(str1) > len(str2) {
        commonstr = str2 
        longstr = str1 
    } else {
        commonstr = str1
        longstr = str2
    }


    for i := len(commonstr); i >= 0; i--{
        checkstr := commonstr[0:i]
        is_common := true
        fmt.Fprintf(os.Stderr, "DEBUGPRINT[6]: 1071.greatest-common-divisor-of-strings.go:17: checkstr=%+v\n", checkstr)


        if len(checkstr) == 0 {
            return ""
        }

        if len(commonstr) % len(checkstr) != 0 {
            continue
        }

        if len(longstr) % len(checkstr) != 0 {
            continue
        }

        for j := 1 ; j <= len(commonstr) / len(checkstr); j++ {
            if commonstr[(j-1) *len(checkstr): j * len(checkstr)] != checkstr {
                is_common = false
                break
            }
        }

        for j := 1 ; j <= len(longstr) / len(checkstr); j++ {
            if longstr[(j-1) *len(checkstr): j * len(checkstr)] != checkstr {
                is_common = false
                break
            }
        }

        if is_common{
            return checkstr
        }
    }

    return ""
    
}
// @leet end
