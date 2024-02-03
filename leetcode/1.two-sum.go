package main

// @leet start
func twoSum(nums []int, target int) []int {
    mapping := map[int]int{}

    for i, num := range nums{
        _, ok := mapping[num]
        if !ok{
            mapping[num] = i
        }
        
        if val, ok:= mapping[target-num]; ok && mapping[target-num] != i{
           return []int{val, i}
        }

    }

    return []int{}
}
// @leet end

func main(){}
