package main

// @leet start
func findOrder(numCourses int, prerequisites [][]int) []int {
    ans := make([]int, 0)
    status := make([]int, numCourses)
    vertices := make([][]int, numCourses)

    for _, val := range prerequisites{
        vertices[val[1]] = append(vertices[val[1]], val[0])
    }

    var dfs func(u int) bool

    dfs = func(u int)bool {
        status[u] = 1
        for _, v := range vertices[u]{
            if status[v] == 1 {
                return false
            }

            if status[v] == 0 {
                if !dfs(v){
                    return false
                }
            }

        }

        status[u] = 2
        ans = append(ans, u)
        return true
    }


    for i := 0; i< numCourses; i++{
        if status[i]  == 0 {
            if !dfs(i){
                return []int{}
            }
        }
    }
    for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return ans
    
}
// @leet end
