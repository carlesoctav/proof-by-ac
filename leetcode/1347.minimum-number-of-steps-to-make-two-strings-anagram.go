package main

// @leet start
func minSteps(s string, t string) int {
    counter_s := make([]int, 26)
    counter_t := make([]int, 26)

    for _, val := range s {
        counter_s[val - 'a']+=1

    }
    for _, val := range t {
        counter_t[val - 'a']+=1
    }

    ans := 0
    for i, _ := range counter_s {
        if counter_s[i] < counter_t[i]{
            ans += -(counter_s[i] - counter_t[i])
        } else if counter_s[i] > counter_t[i]{
            ans += (counter_s[i] - counter_t[i])
        }
    }

    return ans/2
}
// @leet end
