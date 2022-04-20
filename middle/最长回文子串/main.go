package main

import "fmt"

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen == 1 || strLen == 0 {
		return s
	}
	if strLen == 2 {
		if s[0:1] == s[1:2] {
			return s
		} else {
			return s[0:1]
		}
	}
	max := 0
	res := ""
	for i := 1; i < strLen-1; i++ {
		str1, len1 := asMiddle(s, i, strLen)
		if len1 > max {
			max = len1
			res = str1
		}
		//fmt.Printf("str1: %v|len1: %v\n", str1, len1)
		str2, len2 := asMiddleLeft(s, i, strLen)
		if len2 > max {
			max = len2
			res = str2
		}
		//fmt.Printf("str2: %v|len2: %v\n", str2, len2)
		str3, len3 := asMiddleRight(s, i, strLen)
		if len3 > max {
			max = len3
			res = str3
		}
		//fmt.Printf("str3: %v|len3: %v\n", str3, len3)
	}
	return res
}

func asMiddle(s string, index int, len int) (string, int) {
	left := index
	right := index
	return getLongestStr(left, right, s, len)
}

func asMiddleLeft(s string, index int, len int) (string, int) {
	left := index
	right := index + 1
	return getLongestStr(left, right, s, len)
}

func asMiddleRight(s string, index int, len int) (string, int) {
	left := index - 1
	right := index
	return getLongestStr(left, right, s, len)
}

func getLongestStr(left int, right int, s string, len int) (string, int) {
	for left >= 0 && right < len {
		if s[left:left+1] == s[right:right+1] {
			left -= 1
			right += 1
		} else {
			break
		}
	}
	//fmt.Printf("left: %v|right: %v\n", left, right)
	return s[left+1 : right], right - left - 1
}

func main() {
	str := "cbbd"
	res := longestPalindrome(str)
	fmt.Printf("res: %v\n", res)
}
