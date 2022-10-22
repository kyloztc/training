package main

func maxPower(s string) int {
	if s == "" {
		return 0
	}
	_char := s[0]
	count := 1
	_max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == _char {
			count += 1
		} else {
			count = 1
			_char = s[i]
		}
		_max = max(_max, count)
	}
	if _max == 1 {
		return 0
	}
	return _max
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	maxPower("1231231")
}
