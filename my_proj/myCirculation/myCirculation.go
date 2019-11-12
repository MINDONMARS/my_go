package main

import "fmt"

func main() {
	//输出1到100素数
	//var A, a int
	//
	//A = 1
	//
	//c_point:for A < 100 {
	//			A++
	//
	//
	//
	//			for a = 2; a < A; a++ {
	//
	//				if A % a == 0{
	//					goto c_point
	//				}
	//			}
	//
	//			fmt.Printf("%d\n", A)
	//}
	//99乘法表
	for i:=1; i<=9; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}
}
