package main

import "sort"

// @leet start
func sortPeople(names []string, heights []int) []string {
    type tuple struct {
        F string
        S int
    } 
    names_and_heights := make([]tuple, 0)

    for i:= 0; i< len(names);i++{
        names_and_heights = append(names_and_heights, tuple{names[i], heights[i]})
    }

    less := func(i, j int) bool{
        return names_and_heights[i].S < names_and_heights[j].S

    }
    sort.Slice(names_and_heights, less)
    answer := make([]string, 0)
    for _, nm := range names_and_heights{
        answer = append(answer, nm.F)
    }
    return answer
}
// @leet end 
