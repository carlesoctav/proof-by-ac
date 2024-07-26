package main

// @leet start
func findRestaurant(list1 []string, list2 []string) []string {
    
    minSum := int(1e9)
    ans := make([]string, 0)
    for i, v := range list1{
        for j, v2 := range list2{
            if v == v2{
                if i+j < minSum{
                    ans = make([]string, 0)
                    ans = append(ans, list1[i])
                    minSum = i+j
                } else if i+j == minSum{
                    ans = append(ans, list1[i])
                }
            }
        }
    }
    return ans 
}
// @leet end
