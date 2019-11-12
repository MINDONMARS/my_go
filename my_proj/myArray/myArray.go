package main

import "fmt"

func main() {
	//var n [10] int
	//var i, j int
	//
	////为数组初始化元素
	//for i=0; i < 10; i++ {
	//	n[i] = i + 100
	//}
	//
	//for j=0; j < 10; j+=1 {
	//	fmt.Printf("element[%d] = %d\n", j, n[j])
	//}

	//多维数组
	//var threedim [5][10][4] int
	//
	//fmt.Println(threedim)
	var a = [3][5]int{{1, 2, 3, 4, 5}, {0, 9, 8, 7, 6}, {3, 4, 5, 6, 7}}
	var i, j int
	for i=0; i < 3; i++ {
		for j=0; j < 5; j++{
			fmt.Printf("a[%d][%d] = %d ", i, j, a[i][j])
		}
		fmt.Println()
	}

}