package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	if left == right {
		return 0
	}
	max := 0
	for left < right {
		min, leftMin := getMin(left, right, height)
		area := calculateArea(left, right, min)
		if area > max {
			max = area
		}
		if leftMin {
			left += 1
		} else {
			right -= 1
		}
	}
	return max
}

func getMin(left int, right int, height []int) (int, bool) {
	leftMin := true
	min := height[left]
	if min > height[right] {
		min = height[right]
		leftMin = false
	}
	return min, leftMin
}

func calculateArea(left int, right int, height int) int {
	return (right - left) * height
}

func main() {
	height := []int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	}
	res := maxArea(height)
	fmt.Printf("res: %v\n", res)
}
