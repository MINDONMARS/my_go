package main

import "fmt"

//信道可以做生成器使用，作为一个特殊的例子，它还可以用作随机数生成器。如下是一个随机01生成器:

func rand01() chan int { //0，1生成器
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}

	}()

	return ch
}

func main() {
	generator := rand01()

	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}
