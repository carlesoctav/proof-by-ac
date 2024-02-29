package main

// @leet start
func checkPalindrome(s string) bool {

    if len(s) == 0 || len(s) == 1{
        return true
    }

    return s[0] == s[len(s)-1] && checkPalindrome(s[1:len(s)-1])
}
func firstPalindrome(words []string) string {

    for _, value := range words{
        isPalindrome := checkPalindrome(value)

        if isPalindrome{
            return value
        }
    }
    return ""
}
// @leet end
