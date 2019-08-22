package main

import "fmt"

/**
两种切片方式的比较：
基于数组创建切片：直接引用数组，这个数组是事先存在的，程序员是可见的
make直接创建切片：通过mak来创建切片，make也会创建一个数组，是由切片在底层进行维护的，程序员是看不见的。
*/
func main() {

	slice1()
	fmt.Println()
	slice2()
}

/**
基于数组创建数组切片

可以看到，所谓切片就是对数据进行切割，取出我们想要的数组片段。
// 得到所有的数组切片：
mySlice = myArray[:]
// 得到所有的数组前 n 个切片：
mySlice bi= myArray[:n]
// 得到所有的数组后 m 个切片：
mySlice = myArray[m:]
// 得到所有的数组从 n 到 m 个切片：
mySlice = myArray[n:m]
特别说明：这里的切片是看数组的个数，而非下标。且是左开右闭区间
*/
func slice1() {
	// 定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组创建一个切片
	var mySlice []int = myArray[2:5]

	var mySlice2 []int = myArray[5:8]

	fmt.Println("Element of myArray :")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	// 我们在看看我们的切片
	fmt.Println("\nElement of Slice : ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// 切片中，GO内置了几个函数：len() 查看当前切片长度, cap() 查看当前切片容量, append() 对切片进行追加
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))
	// 这里的 ... 不能丢失，在GO的语义中， ... 相当于一个省略号，在append中，第二个参数起的所有的参数都是待附加的元素，因为mySlice2
	// 中的元素类型为 int ，所以直接传递 mySlice 行不通的，加上省略号相当于把 mySice包含的所有元素打散后传入
	mySlice2 = append(mySlice2, 1, 2, 3)
	mySlice2 = append(mySlice2, mySlice...)
	fmt.Println("\nElement of Slice2 : ")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElement of Slice1 : ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	// 讲mySlice中的元素复制到mySlice2中去，注意：复制的值为最短的切片长度为复制的元素个数，
	// 即 n = min(len(mySlice2) , len(mySlice))为我们的复制的前 n 个元素
	copy(mySlice2, mySlice)
	fmt.Println("\n将mySlice的前 n 个元素复制到 mySlice2中后 ， mySlice2中的数据 : ")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
}

/**
直接创建切片
通过make方式创建切片可以指定切片大小和容量
如果没有给切片的各个元素赋值，便会使用默认值
通过make方式创建切片对应的数组是由make底层维护的，对外不可见，即只能通过slice去访问各个元素
*/
func slice2() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice1 := make([]int, 5)
	mySlice1[1] = myArray[0]

	mySlice2 := make([]int, 5, 10)
	mySlice2[1] = myArray[0]

	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
}
