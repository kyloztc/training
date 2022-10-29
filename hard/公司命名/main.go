package main

func distinctNames(ideas []string) int64 {
	ans := 0
	group := make([]map[string]bool, 26)
	for i, _ := range group {
		group[i] = make(map[string]bool)
	}
	for _, idea := range ideas {
		group[idea[0]-'a'][idea[1:]] = true
	}
	for i, a := range group {
		for _, b := range group[:i] {
			m := 0
			for k, _ := range a {
				if b[k] {
					m += 1
				}
			}
			ans += (len(a) - m) * (len(b) - m)
		}
	}
	return int64(ans) * 2
}
