package main

import "fmt"

/**
 * 看看GO语言语言中函数的使用：
 * 函数的使用在之前就已经用过了，但是呢，还是只简单的用法，GO语言也支持不定参数的传入，即：func myFunc1(args ... int){}
 * 这里的(args ... int)便是我们传入的不定参数。需要注意的是，和Java,Scala一样，我们的不定参数需要在所有的参数的后面，很好理解，因为我们
 * 不知道谁是最后一个参数。这也就是我们所谓的语法糖(syntactic sugar)
 * 如果没有这个语法糖，我们就需要在传入参数时设置一个数组切片。
 **/
func main() {
	var v float32 = 1.234
	var v2 int32 = 223
	myFunc4(2, 5, 4, "sun", 2.258, v, v2)
}

func myFunc1(args ...int) {
	// 我们必须要在传入的参数后面加上 ... ，或者会报参数异常，他会要求我们将 ...int 改成 ...[]int
	myFunc2(args...)
}

func myFunc2(args ...int) {
	// 同养，我们的参数传递支持任意的数组切片
	myFunc3(args[2:]...)
}

func myFunc3(args ...int) {

}

/**
 * 在上面的函数中，我们讨论的是规定类型的不定参数，但是如果我们要是希望传递的是任意类型，可以指定类型为 interface{} ,举个例子：
 * 我们在GO语言中， 有执行过 fmt.Printf()的函数，那么，他的原型就是：
 * func Printf(format string, a ...interface{})
 */
func myFunc4(args ...interface{}) {
	// 对传进来的参数做一个类型匹配
	for _, t := range args {
		switch t.(type) {
		case int:
			fmt.Println(t, "-> int")
		case string:
			fmt.Println(t, "-> string")
		case int64:
			fmt.Println(t, "-> int64")
		case float32:
			fmt.Println(t, "-> float32")
		case float64:
			fmt.Println(t, "-> float64")
		default:
			fmt.Println(t, "-> default")
		}
	}
}
