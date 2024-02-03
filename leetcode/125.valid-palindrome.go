package main

import (
	"fmt"
	"unicode"
)

// @leet start
func palindrome(s string) bool {

    if len(s) ==1 || len(s) ==0{
        return true
    }
    return (s[0] == s[len(s)-1]) && palindrome(s[1:len(s)-1])
}
func isPalindrome(s string) bool {

    removeNonAlphaNumeric := func(str string) string {
        isAlpha := []rune{} 
        for _, char := range []rune(str){
            if unicode.IsLetter(char) || unicode.IsNumber(char){
                char = unicode.ToLower(char)
                isAlpha = append(isAlpha, char)

            }
        }

        return string(isAlpha)

    }
    a:= removeNonAlphaNumeric(s)
    fmt.Println(a)
    return palindrome(removeNonAlphaNumeric(a)) 
}
// @leet end
