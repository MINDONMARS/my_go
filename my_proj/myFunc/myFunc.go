package main

func swap(x *int, y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {

	var (
		a, b int
	)
	a = 1
	b = 2

	swap(&a, &b)
	print(a, b)
	
}
