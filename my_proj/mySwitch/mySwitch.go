package main

import "fmt"

/*
1. 支持多条件匹配

switch{
    case 1,2,3,4:
    default:
}
2. 不同的 case 之间不使用 break 分隔，默认只会执行一个 case。

3. 如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止。

switch{
    case 1:
    ...
    if(...){
        break
    }

    fallthrough // 此时switch(1)会执行case1和case2，但是如果满足if条件，则只执行case1

    case 2:
    ...
    case 3:
}
*/

func main() {
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70: grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是%s\n", grade)

	//type-switch
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型：%T", i)
	case int:
		fmt.Printf("x的类型是int：%T", i)
	case float64:
		fmt.Printf("x的类型是float64：%T", i)
	case func(int)float64:
		fmt.Printf("x的类型是func(int)：%T", i)
	case bool, string:
		fmt.Printf("x是bool型或string型：%T", i)
	default:
		fmt.Printf("未知类型：%T", i)
	}

	//fallthrought
	switch {
	case false:
		fmt.Printf("1, case false\n")
		fallthrough
	case true:
		fmt.Printf("2, case true\n")
		fallthrough
	case false:
		fmt.Printf("3, case false\n")
		fallthrough
	default:
		fmt.Printf("4, default case\n")
	}
}
