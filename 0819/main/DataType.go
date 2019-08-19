package main

import (
	"fmt"
	_ "fmt"
	"strconv"
)

func main(){
	// 基本数据类型和String的切换
	// 第一种方式 ， fmt.Sprintf("%参数" , 表达式)
	var num1 int = 99
	var str string
	str = fmt.Sprintf("%d" , num1)
	fmt.Printf("the type of str is %T and the value of str is %q \n" , str ,str)

	var num2 float64 = 2.568
	str = fmt.Sprintf("%f" , num2)
	fmt.Printf("the type of str is %T and the value of str is %q \n" , str ,str)

	var flag bool = true
	str = fmt.Sprintf("%t" , flag)
	fmt.Printf("the type of str is %T and the value of str is %q \n" , str ,str)

	// 第二种方式：使用strconv函数
	str = strconv.FormatInt(int64(num1) , 10)
	fmt.Printf("the type of str is %T and the value of str is %q \n" , str ,str)
	/**
		float64() : 格式化成64位的float
		'f' : 格式
		10 : 表示小数点保留10位
		64 : 表示这个小数是64位的
	 */
	str = strconv.FormatFloat(float64(num2) , 'f' , 10 , 64)
	fmt.Printf("the type of str is %T and the value of str is %q \n" , str ,str)

	str = strconv.FormatBool(flag)
	fmt.Printf("the type of str is %T and the value of str is %q \n" , str ,str)

	// string转基本数据类型
	var str_flag string = "true"
	/**
	b, _ := strconv.ParseBool(str_flag)
		ParseBool的返回值有两个，一个是当前的转化后的类型，另一个是一个error类型，相当于Java中的异常
	 */
	b, _ := strconv.ParseBool(str_flag)
	fmt.Printf("the type of b is %T and the value of b is %v \n" , b , b)
	var str_int64 string = "6985"

	i, _ := strconv.ParseInt(str_int64, 10, 64)
	fmt.Printf("the type of i is %T and the value of i is %v \n" , i , i)
	// 返回的是int64，如希望得到int32,使用如下操作
	// int32(i)

}
