package main

// @leet start
func partition(s string) [][]string {
    ans := make([][]string, 0)
    curr := make([]string, 0)
    var pdrome func(s string) bool
    var backtracking func(idx int)
    pdrome = func(s string) bool{
        if len(s) == 1{
            return true
        }
        if len(s) == 2{
            return s[0] == s[1]
        }

        return s[0] == s[len(s)-1] && pdrome(s[1:len(s)-1])
    }
    backtracking = func(idx int){
        if idx == len(s){
            ans = append(ans, append([]string{}, curr...))
        }

        for i:= idx ; i < len(s); i++{
            if pdrome(s[idx: i+1]) {
                curr = append(curr, s[idx: i+1])
                backtracking(i+1)
                curr = curr[: len(curr)-1]
            }
        }
    }
    backtracking(0)

    return  ans
}
// @leet end
