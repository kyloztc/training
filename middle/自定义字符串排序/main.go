package main

import (
	"fmt"
	"sort"
)

func customSortString(order string, s string) string {
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}
	strByte := []byte(s)
	sort.Slice(strByte, func(i, j int) bool {
		return orderMap[strByte[i]] <= orderMap[strByte[j]]
	})
	return string(strByte)
}

func main() {
	order := "cbafg"
	s := "abcd"
	res := customSortString(order, s)
	fmt.Printf("res: %v\n", res)
}
