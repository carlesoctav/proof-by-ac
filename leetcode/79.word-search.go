package main

// @leet start
func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    dr := []int{0,0,-1,1}
    dc := []int{1,-1,0,0}

    var dfs func(i, j, lenWord int) bool

    dfs = func(i,j, lenWord int) bool{
        if i < 0 || i>=m || j<0 || j>=n || lenWord == len(word){
            return false
        }

        if board[i][j] == '*' || word[lenWord] != board[i][j]{
            return false
        }

        if lenWord == len(word) -1 {
            return true
        }
        tmp := board[i][j]
        res := false
        board[i][j] = '*'
        for mv := 0; mv<4; mv++{
            newI, newJ := i+dr[mv], j+dc[mv]
            res = res || dfs(newI, newJ, lenWord+1)
        }
        board[i][j] = tmp
        return res
    }
    for i:=0; i<m;i++{
        for j:=0;j<n;j++{
            if dfs(i,j,0){
                return true
            }
        }
    }
    return false
}
// @leet end
