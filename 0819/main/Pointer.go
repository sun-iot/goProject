package main

import "fmt"

func main(){
	var a int
	var ptr *int
	var pptr **int

	a = 300
	// ptr的指针指向a的地址
	ptr = &a

	pptr = &ptr

	// 获取pptr的值
	fmt.Printf("pptr = %d" , **pptr)
}
