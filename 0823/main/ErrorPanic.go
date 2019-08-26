package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	result, err := Operate(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("num1 / num2 = ", result)
	}
}

// 先看下函数错误处理模式的定义
func errorFunc() (i int, err error) {

	if err != nil {
		// 错误处理
		fmt.Println(err)
	} else {
		// 正常处理逻辑
	}
	// 为了方便不报错，返回的err统一设为nil
	return 0, nil
}

// 自定义一个路径错误的error类型
type CustomerDivError struct {
	a    int
	b    int
	Info string
	Err  error
}

// 实现Error方法，这里用到指针的方式,这样的话，我们就可以返回一个PathError了
func (e *CustomerDivError) Error() string {
	return e.Info + ":" + e.Err.Error()
}

// 在下面的代码逻辑中，syscall.Stat_t 失败返回err,将err打包给我们的PathError对象
func Operate(num1 int, num2 int) (result int, err error) {
	data := CustomerDivError{
		a:    num1,
		b:    num2,
		Info: "",
		Err:  nil,
	}
	if num2 == 0 {
		// data.Info = "除数不可为0"
		data.Info = data.Err.Error()
		return
	} else {
		return num1 / num2, nil
	}
}

/**
GO语言实现文件的复制
Copy函数抛出异常，但是我们的dstFile和srcFile依然可以正常关闭
如果一句话清理不干净，我们也可以在defer后面接一个匿名函数
*/
func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
