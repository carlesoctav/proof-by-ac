package main


// @leet start
func equalSubstring(s string, t string, maxCost int) int {
    ans := 0
    prefixSum := make([]int,len(s)+1)
    prefixSum[0] = 0
    for i, _ := range s{
        diff := (int(s[i]) - int(t[i]))
        if diff >=0 {
            prefixSum[i+1] =  diff + prefixSum[i]
        } else {
            prefixSum[i+1] = -diff + prefixSum[i]
        }
    }

    i, j := 0, 0

    for  j != len(s){
        cost := prefixSum[j] - prefixSum[i]
        if cost > maxCost{
            i++
            continue
        }

        ans = max(j-i, ans)
        j++
    }
    return ans
}
// @leet end
