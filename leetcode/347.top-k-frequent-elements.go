package main

import "sort"

// @leet start
func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)

    for _, value := range nums {
        if _, ok := freq[value]; !ok {
            freq[value] = 1
        } else {
            freq[value] += 1
        }
    }
    
    keys := make([]int, 0, len(freq))

    for key := range freq{
        keys = append(keys, key)
    }

    sort.SliceStable(keys, func(i, j int) bool {
        return freq[keys[i]] >freq[keys[j]]
    })


    return keys[:k]
}

// bisa O(n) btw, create a freq table with key is num occurence, and the value
// are the num that occurs key time
// @leet end
