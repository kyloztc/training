package main

import "fmt"

func parseBoolExpr(expression string) bool {
	operatorStack := make([]rune, 0)
	boolStack := make([]rune, 0)
	for _, c := range expression {
		switch c {
		case '!', '&', '|':
			operatorStack = append(operatorStack, c)
		case '(', 'f', 't':
			boolStack = append(boolStack, c)
		case ')':
			boolCount := make(map[rune]int)
			boolCount['t'] = 0
			boolCount['f'] = 0
			for boolStack[len(boolStack)-1] != '(' {
				boolStr := boolStack[len(boolStack)-1]
				boolCount[boolStr] += 1
				boolStack = boolStack[:len(boolStack)-1]
			}
			boolStack = boolStack[:len(boolStack)-1]
			operator := operatorStack[len(operatorStack)-1]
			operatorStack = operatorStack[:len(operatorStack)-1]
			currentBool := 'f'
			falseCount := boolCount['f']
			trueCount := boolCount['t']
			if (operator == '&' && falseCount == 0) || (operator == '|' && trueCount != 0) || (operator == '!' && falseCount != 0) {
				currentBool = 't'
			}
			boolStack = append(boolStack, currentBool)
		}
	}
	return boolStack[0] == 't'
}

func main() {
	res := parseBoolExpr("!(f)")
	fmt.Printf("res: %v\n", res)
}
