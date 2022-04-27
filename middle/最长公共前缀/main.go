package main

import "fmt"

type PrefixTree struct {
	Value  string
	Child  map[string]*PrefixTree
	HasEnd bool
}

func NewPrefixNode(value string) *PrefixTree {
	return &PrefixTree{
		Value: value,
		Child: make(map[string]*PrefixTree),
	}
}

func (t *PrefixTree) AddStr(str string) {
	if len(str) == 0 {
		return
	}
	first := str[0:1]
	end := false
	if len(str) == 1 {
		end = true
	}
	_, exist := t.Child[first]
	if !exist {
		t.Child[first] = NewPrefixNode(first)
	}
	t.Child[first].AddStr(str[1:])
	t.Child[first].SetEnd(end)
	return
}

func (t *PrefixTree) SetEnd(end bool) {
	if t.HasEnd {
		return
	}
	t.HasEnd = end
}

func (t *PrefixTree) GetLongestPrefix() string {
	res := t.Value
	if len(t.Child) != 1 || t.HasEnd {
		return res
	}
	for _, node := range t.Child {
		res += node.GetLongestPrefix()
	}
	return res
}

func longestCommonPrefix(strs []string) string {
	prefixTree := NewPrefixNode("")
	for _, str := range strs {
		if str == "" {
			return ""
		}
		prefixTree.AddStr(str)
	}
	res := prefixTree.GetLongestPrefix()
	return res
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Printf("res: %v\n", res)
}
