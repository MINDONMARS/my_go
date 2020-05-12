package main

import "fmt"

func get_notifitication(user string) chan string {
	/*
	 * 此处可以查询数据库获取消息
	 */
	notifications := make(chan string)

	go func() {
		notifications <- fmt.Sprintf("Hi %s, welcome to weibo.com!", user)
	}()

	return notifications
}

func main() {
	jack := get_notifitication("jack")

	joe := get_notifitication("joe")

	fmt.Println(<-jack)
	fmt.Println(<-joe)
}
