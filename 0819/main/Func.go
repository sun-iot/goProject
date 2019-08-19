package main

import "fmt"

func main(){
	s, i := swap("yang", "sun")
	fmt.Println(s + " " + i )

	// 比较值传递与引用传递
	var a , b int
	a = 10
	b = 50
	// 先看值传递
	fmt.Printf("传递前：\n a = %d , b = %d \n" , a , b )
	//temp := swapValue(a, b)
	//temp := swapReference(&a, &b)
	temp := swapReference2(&a, &b)
	fmt.Printf("传递后：\n a = %d , b = %d \n" , a , b )

	fmt.Println(temp)

}
/**
	交换两个字符串
 */
func swap(str1 , str2 string)(string , string ){
	return str2 , str1
}


/**
	值传递
 */
func swapValue(x , y int)(temp int){
	temp = x
	x = y
	y = temp
	return temp
}
/**
	引用传递
 */
func swapReference(x *int, y *int)(temp int){
	temp = *x
	*x = *y
	*y = temp
	return temp
}
/**
不知道是个啥
 */
func swapReference2(x , y *int)(temp *int){
	*temp = *x
	*x = *y
	*y = *temp
	return temp
}
