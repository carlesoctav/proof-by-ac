package main

// @leet start
func maxNumberOfBalloons(text string) int {
    listChar := make([]int, 26)

    for _, val := range text{
        listChar[val - 'a']++
    }

    count := 0
    ballons := []rune{'a', 'b', 'o', 'l', 'n'}
    substract := []int{1, 1, 2, 2, 1}

    for {
        for i, val := range ballons{
            listChar[val - 'a'] -= substract[i]
            if listChar[val -'a'] < 0{
                return count
            }
        }
        count+=1
    }
    
}
// @leet end
