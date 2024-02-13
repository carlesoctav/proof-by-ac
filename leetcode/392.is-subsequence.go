package main // @leet start
func isSubsequence(s string, t string) bool {
    sPointer := -1
        


    if len(s) ==0{
        return true
    }

    for _, char := range []rune(t){
        if char == rune(s[sPointer]){
            sPointer++
        }

        if sPointer == len(s) -1{
            return true
        }
    }

    return false
    
}
// @leet end
