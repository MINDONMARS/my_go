package main

import "fmt"

func main() {
	/*
	指针使用流程：
	定义指针变量。
	为指针变量赋值。
	访问指针变量中指向地址的值。
	*/
	//var a int = 20
	//var ip *int
	//ip = &a
	//
	//fmt.Printf("变量a的地址是%x\n", &a)
	//
	//fmt.Printf("ip变量存储的指针地址： %x\n", ip)
	//
	////使用指针访问值
	//fmt.Printf("*ip变量的值：%d\n", *ip)
	//
	//var ptr *int
	//
	//fmt.Printf("ptr 的值为：%x\n", ptr)
	//
	//if ptr != nil {
	//	fmt.Printf("ptr 不是空指针")
	//} else {
	//	fmt.Printf("ptr 是空指针")
	//}
	//fmt.Println("\n----------------------------------")
	//
	////指针数组
	//const MAX int = 3
	//
	//aa := []int{10, 100, 200}
	//var i int
	//
	//for i=0; i<MAX; i++ {
	//	fmt.Printf("a[%d] = %d\n", i, aa[i])
	//}
	//
	//var ptrr [MAX] *int
	//
	//aaa := []int{100, 200, 300}
	//var j int
	//for j = 0; j < MAX; j++{
	//	ptrr[j] = &aaa[j]
	//}
	//for j = 0; j < MAX; j++{
	//	fmt.Printf("aaa[%d] = %d\n", j, *ptrr[j])
	//}
	//
	////创建指针数组的时候，不适合用 range 循环。
	//
	//const  MAKS  = 3
	//
	//number := [MAKS]int{5, 6, 7}
	//var ptrrr [MAX]*int
	//for iii := 0; iii < MAKS; iii++ {
	//	ptrrr[iii] = &number[iii]
	//}
	//for iii, xxx :=range ptrrr{
	//	fmt.Printf("指针数组：索引：%d 值%d 值的内存地址：%d\n", iii, *xxx, xxx)
	//}
	//指针作为函数参数
	var a int = 100
	var b int = 200

	fmt.Printf("交换前a的值：%d\n", a)
	fmt.Printf("交换前b的值：%d\n", b)

	swap(&a, &b)

	fmt.Printf("交换后a的值：%d\n", a)
	fmt.Printf("交换后b的值：%d\n", b)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
