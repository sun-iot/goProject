# GO语言数据类型

在GO编程语言中，数据类型用于声明函数和变量；

GO语言按照类型分有如下几种：

| 序号 | 类型和描述                                                   |
| :--: | :----------------------------------------------------------- |
|  1   | **布尔型** 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
|  2   | **数字类型** 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
|  3   | **字符串类型:** 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 |
|  4   | **派生类型:** 包括：(a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型 |

### GO语言数字类型

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **uint8** 无符号 8 位整型 (0 到 255)                         |
| 2    | **uint16** 无符号 16 位整型 (0 到 65535)                     |
| 3    | **uint32** 无符号 32 位整型 (0 到 4294967295)                |
| 4    | **uint64** 无符号 64 位整型 (0 到 18446744073709551615)      |
| 5    | **int8** 有符号 8 位整型 (-128 到 127)                       |
| 6    | **int16** 有符号 16 位整型 (-32768 到 32767)                 |
| 7    | **int32** 有符号 32 位整型 (-2147483648 到 2147483647)       |
| 8    | **int64** 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

**浮点类型**

| 序号 | 类型和描述                        |
| :--- | :-------------------------------- |
| 1    | **float32** IEEE-754 32位浮点型数 |
| 2    | **float64** IEEE-754 64位浮点型数 |
| 3    | **complex64** 32 位实数和虚数     |
| 4    | **complex128** 64 位实数和虚数    |

**其他数字类型**

| 序号 | 类型和描述                               |
| :--- | :--------------------------------------- |
| 1    | **byte** 类似 uint8                      |
| 2    | **rune** 类似 int32                      |
| 3    | **uint** 32 或 64 位                     |
| 4    | **int** 与 uint 一样大小                 |
| 5    | **uintptr** 无符号整型，用于存放一个指针 |

# GO语言中变量的声明

> + 指定变量类型，如果没有初始化，则变量默认为零值
> + 根据值自行判定变量类型
> + 省略var，注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
>
> ```go
> v_name = value 
> 
> var intval int 
> intval := 1 // 报错，左侧未声明新的变量
> intval , intval_new := 1 , 2 // 可以编译，声明了新的变量
> 
> ```
>
> + 对于多变量的声明，可以使用因式分解关键字的写法，如下所示：
>
> ```go
> var (
> 	vname1 v_type1
>     vname2 v_type2
>     ......
> )
> ```

## 数据类型和string的转换

**数据类型 -> string**

```go
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
```

**string -> 基本数据类型**

```go
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
```

> 注意：再将string类型转成基本数据类型时，要确保String类型能转化成有效的数据，对于没有转成功的，会给我们的接收的值默认值。

# GO语言的数组

```go
	// 数组的定义，与C语言相似
	var array_int [10]int // 定义一个长度为10的int数组
	var i , j int

	for i = 0 ; i < 10 ; i ++ {
		array_int[i] = i * i
	}

	for j = 0 ; j < 10 ; j ++ {
		fmt.Printf("Element[%d] = %d \n" , j , array_int[j])
	}

	// 对于根据数组长度进行自定义数据长度的数据，我们定义如下
	var array_str  = [...]string{"sun" , "yang" , "hello" , "corp" ,"ci" ,"www"}
	for i = 0 ; i < len(array_str); i ++{
		fmt.Println(array_str[i])
	}
```

**函数中传入数组**

```go
func getAvg(arr [10]int) (avg float32) {
	var i , sum int

	for i=0 ; i < len(arr) ; i ++{
		sum += arr[i]
	}
	avg = float32(sum /len(arr))
	return avg
}
```

# GO语言select语句

> 说明：select语句是GO通信中的一个控制结构，类似于通信的switch语句，每一个case必须是一个通信操作，要么发送要么接收；
>
> select是随机执行的一个可执行的case，如果没有case，它将阻塞，直到有case可运行。

```go
select {
    case communication clause :
    statement(s);
    case communication clause:
    statement(s);
    
	default : /*可选*/
    statement(s);
}
```

# GO语言指针

> 一个指针变量指向一个值得内存地址

```go
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
```

```go
// 对于上面的程序，很显然，第三种方式会抱一个错误：
传递前：
 a = 10 , b = 50 
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x0 pc=0x492eec]

goroutine 1 [running]:
main.swapReference2(...)
	D:/corp-ci-go_ide/demo/0819/main/Func.go:53
main.main()
	D:/corp-ci-go_ide/demo/0819/main/Func.go:17 +0x1dc

Process finished with exit code 2
// 无效的内存地址或无指针取消引用
```

## GO指向指针的指针

> 如果一个指针变量存放的是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量
>
> 当定义一个指向指针的指针变量时，第一个指针指向第二个指针地址，第二个指针存放变量的地址

```go
	var a int
	var ptr *int
	var pptr **int
	a = 300 
	// ptr的指针指向a的地址
	ptr = &a
	pptr = &ptr
	// 获取pptr的值
	fmt.Printf("pptr = %d" , **pptr)
```

# GO语言结构体

> GO语言的结构体，和我们的C语言中的结构体是一样的，同养，可以理解为我们Java中的类。其基础定义如下：
>
> ```go
> type struct_type struce{
>     momber definition ;
>     ......
> }
> ```

```go
type Books struct {
	title string
	author string
	subject string
	book_id int
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
	// 指针传递结构体
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
```

# GO语言切片（Slice）

> GO语言切片是对数组的抽象，简单的理解：切片是一个动态数组

**对于切片的定义**

> 1. 可以声明一个未指定大小的数组来定义切片，切片不需要指定长度
> 2. 使用make()函数来创建切片