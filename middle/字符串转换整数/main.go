package main

import (
	"fmt"
	"math"
)

type automaton struct {
	Ans int
	Sign int
	State string
	Table map[string][]string
}

func NewAutomaton() *automaton {
	table := map[string][]string{
		"start": {"start", "signed", "in_number", "end"},
		"signed": {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end": {"end", "end", "end", "end"},
	}
	state := "start"
	return &automaton{
		Ans:   0,
		Sign:  1,
		State: state,
		Table: table,
	}
}

func (a *automaton) get(s string) {
	colIndex := a.getStateCol(s)
	state := a.Table[a.State][colIndex]
	a.State = state
	if state == "in_number" {
		//fmt.Printf("s: %v|int: %v\n", s, int(s[0] - 48))
		a.Ans = a.Ans * 10 + int(s[0] - 48)
		if a.Sign == 1 && a.Ans > math.MaxInt32 {
			a.Ans = math.MaxInt32
		}
		if a.Sign == -1 && a.Ans > math.MaxInt32+1 {
			a.Ans = math.MaxInt32+1
		}
	}
	if state == "signed" {
		if s == "-" {
			a.Sign = -1
		}
	}
}

func (a *automaton) getStateCol(s string) int {
	switch s {
	case " ":
		return 0
	case "+", "-":
		return 1
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return 2
	default:
		return 3
	}
}

func myAtoi(s string) int {
	auto := NewAutomaton()
	for i := 0; i < len(s); i++ {
		auto.get(s[i:i+1])
	}
	return auto.Ans * auto.Sign
}

func main() {
	str := "   -42"
	res := myAtoi(str)
	fmt.Printf("res: %v\n", res)
}
