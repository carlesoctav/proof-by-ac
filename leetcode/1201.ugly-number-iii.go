package main

// @leet start
func nthUglyNumber(n int, a int, b int, c int) int {
    ans := -1
    lo, hi := 1, int(1e10)
    var gcd func(a,b int)int
    gcd = func(a,b int) int {
        if a == 0{
            return b
        }
        return gcd(b % a, a)
    }
    ab := a*b / gcd(a,b)
    ac := a*c / gcd(a,c)
    bc := b*c / gcd(b,c)
    abc := ab*c / gcd(ab,c)
    condition := func(mid int) bool {
        total := mid / a + mid /b + mid /c - mid/(ab) - mid/(ac) - mid/(bc) + mid/(abc)
        return total >= n
    }

    for  lo<=hi {
        mid := lo + (hi-lo)/2
        if condition(mid){
            hi = mid -1
            ans = mid
        } else {
            lo = mid+1
        }
    }
    return ans 
    
}
// @leet end
