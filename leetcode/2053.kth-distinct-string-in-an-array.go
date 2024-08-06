package main

// @leet start
func kthDistinct(arr []string, k int) string {
    counter := make([]int, 0, 0)
    counter_int := 0
    mapping := make(map[string]int)

    for _, val:=  range arr {
        if s, ok := mapping[val]; !ok {
            mapping[val] = counter_int
            counter = append(counter, 1)
            counter_int++
        } else {
            counter[s]++
        }
    }

    another_counter := 0
    for i, val := range counter {
        if val == 1{
            another_counter++
        }

        if another_counter == k {
            for k, v:= range mapping {
                if v == i{
                    return k
                }
            }
        }
    }

    return ""
    
}
// @leet end
