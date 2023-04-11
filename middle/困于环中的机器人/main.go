package main

func isRobotBounded(instructions string) bool {
	x := 0
	y := 0
	// 0 北 1东 2南 3西
	direction := make([][]int, 4)
	direction[0] = []int{3, 1}
	direction[1] = []int{0, 2}
	direction[2] = []int{1, 3}
	direction[3] = []int{2, 0}
	walk := make([][]int, 4)
	walk[0] = []int{0, 1}
	walk[1] = []int{1, 0}
	walk[2] = []int{0, -1}
	walk[3] = []int{-1, 0}
	face := 0
	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case 'G':
			x += walk[face][0]
			y += walk[face][1]
		case 'L':
			face = direction[face][0]
		case 'R':
			face = direction[face][1]
		}
	}
	if face != 0 || (x == 0 && y == 0) {
		return true
	}
	return false
}
