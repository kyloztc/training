package main

import (
	"fmt"
	"strconv"
)

func solveEquation(equation string) string {
	operator := 1
	xCount := 0
	numCount := 0
	strLen := len(equation)
	i := 0
	end := false
	for i < strLen && !end {
		switch string(equation[i]) {
		case "+":
			operator = 1
		case "-":
			operator = -1
		case "x":
			xCount += operator
		case "=":
			end = true
		default:
			_num, _ := strconv.Atoi(string(equation[i]))
			numCount += operator * _num
		}
		i++
	}
	operator = -1
	for i < strLen {
		switch string(equation[i]) {
		case "+":
			operator = -1
		case "-":
			operator = 1
		case "x":
			xCount += operator
		default:
			_num, _ := strconv.Atoi(string(equation[i]))
			numCount += operator * _num
		}
		i++
	}
	fmt.Printf("numCount: %v|x count: %v\n", numCount, xCount)
	if numCount == 0 && xCount == 0 {
		return "Infinite solutions"
	}
	if numCount == 0 {
		return "x=0"
	}
	return "x=" + strconv.Itoa(-numCount/xCount)
}

func main() {
	res := solveEquation("2x=x")
	fmt.Printf("res: %v\n", res)
}
