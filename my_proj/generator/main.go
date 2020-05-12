package main

import "fmt"

func xrange() chan int { // xrange生成自增整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		for i := 0; ; i++ {
			ch <- i // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

func main() {
	generator := xrange()

	for i := 0; i < 1000; i++ {
		fmt.Println(<-generator)
	}
}
