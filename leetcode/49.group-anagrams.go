package main

// @leet start
func groupAnagrams(strs []string) [][]string {
    type slicechar [26]int
    type group map[slicechar][]string

    grouping := make(group)
    get_group_idx := func(str string) slicechar {
        var local_id slicechar 

        for _, value := range []rune(str){
            local_id[int(value) - int('a')] +=1  
        }
        return local_id
    }
    for _, str := range strs{
        var grouping_idx slicechar = get_group_idx(str)
        grouping[grouping_idx] = append(grouping[grouping_idx], str)
    }
    solution := [][]string{}
    for _, value := range grouping {
        solution = append(solution, value)
    }

    return solution

}
// @leet end
