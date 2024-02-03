package main

// @leet start
func containsDuplicate(nums []int) bool {
    check_double := make(map[int]int)
    for _, num := range nums {
        _, ok := check_double[num]
        if !ok {
            check_double[num] = 1
        } else {
            check_double[num]+=1
            if check_double[num] >1{
                return true
            }
        }
    }

    return false

    
}
// @leet end

func main()
