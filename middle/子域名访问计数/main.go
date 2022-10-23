package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	countMap := make(map[string]int)
	for _, domain := range cpdomains {
		count, _ := strconv.Atoi(strings.Split(domain, " ")[0])
		domainList := strings.Split(strings.Split(domain, " ")[1], ".")
		_domain := ""
		for i := len(domainList) - 1; i >= 0; i-- {
			if _domain == "" {
				_domain = domainList[i]
			} else {
				_domain = fmt.Sprintf("%s.%s", domainList[i], _domain)
			}
			_, ok := countMap[_domain]
			if !ok {
				countMap[_domain] = 0
			}
			countMap[_domain] += count
		}
	}
	res := make([]string, 0)
	for k, v := range countMap {
		res = append(res, fmt.Sprintf("%v %v", v, k))
	}
	return res
}

func main() {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	res := subdomainVisits(cpdomains)
	fmt.Printf("res: %v\n", res)
}
