package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main(){

	getStructure()

	getStructureSingle()
}

func getStructure(){
	// 创建一个新的结构体
	fmt.Println(Books{"GO 语言开发" , "www.runoob.com","GO 语言教程" , 65326489})
	// 使用kv键值对
	fmt.Println(Books{title:"GO 语言开发" , author:"www.runoob.com",subject:"GO 语言教程" , book_id:123456})
	// 忽略的字段为0 或空
	fmt.Println(Books{title:"GO 语言开发" , author:"www.runoob.com"})
}

func getStructureSingle(){
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	printBook(Book1)

	printBookPionter(&Book1)
}
/**
	结构体的函数调用
 */
func printBook( book Books ) {
	fmt.Println("=========================================")
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}
/**
	结构体指针
 */
func printBookPionter(book *Books){
	fmt.Println("=========================================")
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}
