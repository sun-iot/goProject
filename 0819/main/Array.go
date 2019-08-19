package main

import (
	"fmt"
)

func main(){
	// 数组的定义，与C语言相似
	var arrayInt [10]int // 定义一个长度为10的int数组
	var i , j int

	for i = 0 ; i < 10 ; i ++ {
		arrayInt[i] = i * i
	}

	for j = 0 ; j < 10 ; j ++ {
		fmt.Printf("Element[%d] = %d \n" , j , arrayInt[j])
	}

	// 对于根据数组长度进行自定义数据长度的数据，我们定义如下
	var arrayStr = [...]string{"sun" , "yang" , "hello" , "corp" ,"ci" ,"www"}
	for i = 0 ; i < len(arrayStr); i ++{
		fmt.Println(arrayStr[i])
	}
	avg := getAvg(arrayInt)
	fmt.Println(avg)
}

func getAvg(arr [10]int) (avg float32) {
	var i , sum int

	for i=0 ; i < len(arr) ; i ++{
		sum += arr[i]
	}
	avg = float32(sum /len(arr))
	return avg
}
