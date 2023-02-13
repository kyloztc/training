package main

func minCostToMoveChips(position []int) int {
	single := 0
	double := 0
	for i := 0; i < len(position); i ++ {
		if position[i] % 2 == 1 {
			single += 1
		} else {
			double += 1
		}
	}
	return min(single, double)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}