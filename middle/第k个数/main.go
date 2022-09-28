package main

import "fmt"

func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}
	kList := make([]int, k)
	kList[0] = 1
	var p3, p5, p7 int
	for i := 1; i < k; i++ {
		x3, x5, x7 := kList[p3]*3, kList[p5]*5, kList[p7]*7
		_min := findMin(x3, x5, x7)
		kList[i] = _min
		if x3 == _min {
			p3++
		}
		if x5 == _min {
			p5++
		}
		if x7 == _min {
			p7++
		}
	}
	fmt.Printf("list: %v\n", kList)
	return kList[k-1]
}

func findMin(p3, p5, p7 int) int {
	_min := p3
	if p5 < _min {
		_min = p5
	}
	if p7 < _min {
		_min = p7
	}
	return _min
}

func main() {
	res := getKthMagicNumber(8)
	fmt.Printf("res: %v\n", res)
}
