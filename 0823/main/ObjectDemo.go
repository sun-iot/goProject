package main

import "fmt"

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	if IntegerLess(a, 3) {
		fmt.Println(a, "Less 3")
	}
}

type Integer int

/**
定义一个新的类型Integer，他和int没有本质区别，只是他为内置的int类型增加了个新的方法Less(),这样实现了Integer后，就可以让整型想一个
普通的类一样使用
*/
// 面向对象的实现
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 面向过程的实现
func IntegerLess(a Integer, b Integer) bool {
	return a < b
}
