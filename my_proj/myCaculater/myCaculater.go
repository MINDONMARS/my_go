package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行：c=%d\n", c)
	c = a - b
	fmt.Printf("第二行：c=%d\n", c)
	c = a * b
	fmt.Printf("第三行：c=%d\n", c)
	c = a / b
	fmt.Printf("第四行：c=%d\n", c)
	c = a % b
	fmt.Printf("第五行：c=%d\n", c)
	a++
	fmt.Printf("第六行：a=%d\n", a)
	b--
	fmt.Printf("第七行：b=%d\n", b)
}
