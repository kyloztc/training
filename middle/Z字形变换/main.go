package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strList := make([]string, 0)
	verticalLen := numRows
	obliqueLen := numRows - 2
	strLen := len(s)
	for i := 0; i < len(s); i += verticalLen + obliqueLen {
		verticalStr := getSubStr(i, verticalLen, s, strLen)
		obliqueStr := getSubStr(i+verticalLen, obliqueLen, s, strLen)
		strList = append(strList, verticalStr)
		strList = append(strList, obliqueStr)
	}
	res := ""
	for i := 0; i < numRows; i++ {
		for index, str := range strList {
			if index%2 == 0 {
				res += getIndexStr(str, i)
			} else {
				oIndex := verticalLen - i - 2
				res += getIndexStr(str, oIndex)
			}
		}
		//fmt.Printf("res: %v\n", res)
	}
	return res
}

func getIndexStr(s string, index int) string {
	if index >= len(s) || index < 0 || s == "" {
		return ""
	}
	return s[index : index+1]
}

func getSubStr(start int, step int, s string, len int) string {
	if start >= len {
		return ""
	}
	if start+step <= len {
		return s[start : start+step]
	} else {
		return s[start:]
	}
}

func main() {
	str := "PAYPALISHIRING"
	numRows := 2
	res := convert(str, numRows)
	fmt.Printf("res: %v\n", res)
}
