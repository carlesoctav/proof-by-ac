package main

// @leet start
func canFinish(numCourses int, prerequisites [][]int) bool {
    vertices := make([][]int, numCourses)
    status := make([]int, numCourses)
    for _, val := range prerequisites{
        vertices[val[1]] = append(vertices[val[1]], val[0])
    }

    var dfs func(u int)bool

    dfs = func(u int) bool{
        status[u] = 1


        for _, v := range vertices[u]{
            if status[v] == 1{
               return false 
            }


            if status[v] == 0{
                if !dfs(v){
                    return false
                }
            }
        }

        status[u] = 2
        return  true
    }

    for i:=0; i <numCourses; i++{
        if status[i] == 0 {
            if !dfs(i){
                return false
            }
        }
    }
    return true
}
// @leet end
