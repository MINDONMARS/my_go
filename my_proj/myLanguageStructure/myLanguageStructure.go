package main

import "fmt"

/*语言结构体*/

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	////创建一个新结构体
	//fmt.Println(Books{"Go语言", "www.runoob,com", "Go语言教程", 54545})
	////忽略的字段为0或空
	//fmt.Println(Books{title: "Go语言", author: "www.runoob.com"})
	//访问结构体成员
	var Book1 Books
	var Book2 Books

	/* book1描述 */
	Book1.title = "Go语言"
	Book1.author = "www.runoob.com"
	Book1.book_id = 656565

	/* book2描述 */
	Book2.title = "python"
	Book2.author = "www.runoob.com"
	Book2.subject = "python 语言教程"
	Book2.book_id = 7898478

	fmt.Println(Book1.title)
	fmt.Println(Book1.author)
	fmt.Println(Book1.subject)
	fmt.Println(Book1.book_id)

	fmt.Println(Book2.title)
	fmt.Println(Book2.author)
	fmt.Println(Book2.subject)
	fmt.Println(Book2.book_id)

	//结构体作为函数参数
	printBook(Book1)
	printBook(Book2)

	//结构体指针
	//var struct_pointer *Books
	//
	//struct_pointer = &Book1
	//
	//fmt.Println(struct_pointer.title)
	var Book3 Books
	Book3.title = "book3"
	Book3.author = "author"
	Book3.book_id = 1
	changeBook(&Book3)
	fmt.Println(Book3)



}

type Rect struct{   //定义矩形类
	x,y float64       //类型只包含属性，并没有方法
	width,height float64
}
func (r *Rect) Area() float64{    //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
	return r.width*r.height         //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}

func changeBook(book *Books){
	book.title = "book1_change"
}

func printBook(book Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject:%s\n", book.subject)
	fmt.Printf("Book book_id:%d\n", book.book_id)
}
