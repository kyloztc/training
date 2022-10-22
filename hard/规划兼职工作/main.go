package main

import (
	"fmt"
	"sort"
)

type Job struct {
	StartTime int
	EndTime   int
	Profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobList := make([]*Job, 0)
	for i := 0; i < len(startTime); i++ {
		jobList = append(jobList, &Job{
			StartTime: startTime[i],
			EndTime:   endTime[i],
			Profit:    profit[i],
		})
	}
	sort.Slice(jobList, func(i, j int) bool {
		return jobList[i].EndTime <= jobList[j].EndTime
	})
	dp := make([]int, len(startTime))
	for i := 0; i < len(startTime); i++ {
		dp[i] = jobList[i].Profit
	}
	for i := 1; i < len(startTime); i++ {
		dp[i] = max(dp[i-1], dp[i])
		for j := i - 1; j >= 0; j-- {
			if jobList[j].EndTime <= jobList[i].StartTime {
				dp[i] = max(dp[i], dp[j]+jobList[i].Profit)
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	startTime := []int{1, 3, 3, 3, 5, 5, 13, 5}
	endTime := []int{17, 14, 8, 11, 14, 7, 17, 9}
	profit := []int{9, 3, 7, 18, 2, 17, 4, 6}
	res := jobScheduling(startTime, endTime, profit)
	fmt.Printf("res: %v\n", res)
}
