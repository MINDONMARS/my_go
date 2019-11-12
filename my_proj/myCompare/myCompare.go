package main

import "fmt"

func main() {
	var a int
	var b int

	a = 21
	b = 10

	if a == b {
		fmt.Printf("第一行：a=b\n")
	} else {
		fmt.Printf("第一行：a!=b\n")
	}
	if a < b {
		fmt.Printf("第二行：a<b\n")
	} else {
		fmt.Printf("第二行：a不小于b\n")
	}
	if a > b {
		fmt.Printf("第三行: a大于b\n")
	} else {
		fmt.Printf("第三行：a不大于b\n")
	}

	a = 5
	b = 20

	if a <= b {
		fmt.Printf("第四行：a小于等于b\n")
	} else {
		fmt.Printf("第四行：a大于b\n")
	}

	if a >= b {
		fmt.Printf("第五行：a>=b\n")
	} else {
		fmt.Printf("a<b")
	}

}
