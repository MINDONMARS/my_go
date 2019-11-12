package main

import (
	"../myMath"
	"fmt"
)

var kid_age int = 10

var x, y int

var (
	a int
	b bool
)

var c, d int = 1, 2

var e, f = 123, "hello"

func addAge()  {
	fmt.Println(mathClass.Add(kid_age, 10))
}

func main() {
	fmt.Println("hello world")

	fmt.Println(mathClass.Add(1,1))

	fmt.Println(mathClass.Sub(1, 2))

	fmt.Println("google" + "baidu")

	addAge()

	g, h := 123, "hello"

	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	_, numb, strs := numbers()

	fmt.Println(numb, strs)
}

func numbers()(int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}