package main

import (
	"fmt"
	"unsafe"
)

//矩形面积
func RectangleArea() int{
	const LENGTH int = 10

	const WIDTH int = 5

	var area int

	area = LENGTH * WIDTH

	return area
}

//枚举常量
func EnumrationConstant() (int, int, int) {
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	return Unknown, Female, Male
}

func EnumrationConstant2() (string, int, uintptr) {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	return a, b, c
}

//特殊常量iota1
func specialConstant1() (int, int, int) {

	const (
		a = iota
		b
		c
	)
	return a, b, c
}

//特殊常量iota2
func specialConstant2() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a,b,c,d,e,f,g,h,i)
}

//特殊常量3
func specialConstant3() {
	const (
		i = 1<<iota
		j = 3<<iota
		k
		l
	)

	fmt.Println(i, j,k,l)
}

func main()  {
	//矩形面积
	fmt.Printf("面积是：%d", RectangleArea())
	//性别枚举
	un, female, male := EnumrationConstant()
	fmt.Println()
	fmt.Println(un, female, male)
	fmt.Println()
	//枚举包含内置函数
	d, e, f := EnumrationConstant2()
	fmt.Println(d, e, f)
	//特殊常量1
	fmt.Println()
	special_a, special_b, special_c := specialConstant1()
	fmt.Println(special_a, special_b, special_c)
	//特殊常量2
	specialConstant2()
	//特殊常量3
	specialConstant3()
}
