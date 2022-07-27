package main

import "fmt"

func trap(height []int) int {
	maxIndex := 0
	max := 0
	for index, h := range height {
		if h > max {
			max = h
			maxIndex = index
		}
	}
	fmt.Printf("max: %v\n", maxIndex)
	res := 0
	res += findLeftSubTrap(height, 0, 1, maxIndex)
	res += findRightSubTrap(height, len(height)-2, len(height)-1, maxIndex)
	return res
}

func findLeftSubTrap(height []int, left int, right int, maxIndex int) int {
	midArea := 0
	sum := 0
	for right <= maxIndex {
		if height[right] < height[left] {
			midArea += height[right]
		} else {
			fmt.Printf("left: %v|right: %v|mid: %v\n", left, right, midArea)
			subSum := (right-left-1)*height[left] - midArea
			midArea = 0
			left = right
			sum += subSum
		}
		right += 1
	}
	fmt.Printf("left: %v\n", sum)
	return sum
}

func findRightSubTrap(height []int, left int, right int, maxIndex int) int {
	midArea := 0
	sum := 0
	for left >= maxIndex {
		if height[left] < height[right] {
			midArea += height[left]
		} else {
			fmt.Printf("left: %v|right: %v|mid: %v\n", left, right, midArea)
			subSum := (right-left-1)*height[right] - midArea
			midArea = 0
			right = left
			sum += subSum
		}
		left -= 1
	}
	fmt.Printf("right: %v\n", sum)
	return sum
}

func main() {
	height := []int{
		4, 2, 0, 3, 2, 5,
	}
	res := trap(height)
	fmt.Printf("res: %v\n", res)
}
