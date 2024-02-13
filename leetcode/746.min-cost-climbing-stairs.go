package main

// @leet start
func minCostClimbingStairs(cost []int) int {
    n := len(cost)+1
    minCostAtStep := make([]int, len(cost)+1)

    minCostAtStep[0] = cost[0]
    minCostAtStep[1] = cost[1]

    for value:= 2; value < n; value++{
        minCostAtStep[value] = min(minCostAtStep[value-1]+ cost[value], minCostAtStep[value-2]+cost[value])
    }
    minCostAtStep[n] = min(minCostAtStep[n-1], minCostAtStep[n-2])
    return minCostAtStep[n]
}
// @leet end
