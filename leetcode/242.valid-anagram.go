package main

// @leet start
func isAnagram(s string, t string) bool {
    counter := map[byte]int{}
    checkCounter := func(counter map[byte]int) bool{
        for _, num := range counter{
           if num != 0{
                return false
            }
        }

        return true

    }
    for i:=0; i< len(s); i++{
        _, ok := counter[s[i]]
        if !ok {
            counter[s[i]] = 1
        } else {
            counter[s[i]] += 1
        }
    }
    for i:=0; i< len(t); i++{
        _, ok := counter[t[i]]
        if !ok {
            counter[t[i]] = 1
        } else {
            counter[t[i]] -= 1
        }
    }

    return checkCounter(counter)

}
// @leet end

