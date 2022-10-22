package main

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	for _, e := range edges {
		if !union(parent, e[0], e[1]) {
			return e
		}
	}
	return nil
}

func find(parent []int, num int) int {
	if parent[num] != num {
		parent[num] = find(parent, parent[num])
	}
	return parent[num]
}

func union(parent []int, from int, to int) bool {
	x, y := find(parent, from), find(parent, to)
	if x == y {
		return false
	}
	parent[x] = y
	return true
}

func main() {

}
