package main

import (
	"fmt"
	"math/rand"
	"time"
)

func do_stuff(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) // 模拟计算
	return 100 - x
}

func branch(x int) chan int { // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
	ch := make(chan int)

	go func() {
		ch <- do_stuff(x)
	}()

	return ch
}

func fanIn(chs ...chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		go func(c chan int) { ch <- <-c }(c) // 复合
	}

	return ch
}

func main() {
	result := fanIn(branch(1), branch(2), branch(3))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
