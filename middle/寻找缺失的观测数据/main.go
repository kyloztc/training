package main

func missingRolls(rolls []int, mean int, n int) []int {
	rollSum := 0
	for _, num := range rolls {
		rollSum += num
	}
	missingSum := (mean * (n + len(rolls))) - rollSum
	if n > missingSum || missingSum > n*6 {
		return []int{}
	}
	missingMean := missingSum / n
	missingMod := missingSum % n
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if i < missingMod {
			res[i] = missingMean + 1
		} else {
			res[i] = missingMean
		}

	}
	return res
}
