package main

func mincostTickets(days []int, costs []int) int {
	memo := make([]int, 366)
	dayMap := make(map[int]bool)
	for _, day := range days {
		dayMap[day] = true
	}
	return dp(1, memo, dayMap, costs)
}

func dp(n int, memo []int, dayMap map[int]bool, cost []int) int {
	if n > 365 {
		return 0
	}
	if memo[n] > 0 {
		return memo[n]
	}
	if dayMap[n] {
		memo[n] = min(min(cost[0]+dp(n+1, memo, dayMap, cost), cost[1]+dp(n+7, memo, dayMap, cost)), cost[2]+dp(n+30, memo, dayMap, cost))

	} else {
		memo[n] = dp(n+1, memo, dayMap, cost)
	}
	return memo[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
