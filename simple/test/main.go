package main

import "fmt"

type A struct {
}

func (*A) Say() {
	fmt.Printf("hello")
}

func SayHi(a *A) {
	a.Say()
}

func main() {
	SayHi(nil)
}
