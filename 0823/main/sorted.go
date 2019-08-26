package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var inFile *string = flag.String("i", "0823/sort/unsorted.txt", "File contains values for sorting")
var outFile *string = flag.String("o", "0823/sort/sorted.txt", "File to receive sorted values")
var algorithm *string = flag.String("a", "bubblesort", "Sort algorithm")

func main() {
	flag.Parse()

	if inFile != nil {
		fmt.Println("inFile", *inFile, "outFile", *outFile, "algorithm", *algorithm)
	}
	values, err := readValues(*inFile)
	if err != nil {
		fmt.Println(err)
	} else {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			quitSort(0, len(values)-1, values...)
		case "bubblesort":
			bubbleSort(values...)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outFile)

	}
}

/**
从文件读取数据
*/
func readValues(inFile string) (values []int, err error) {

	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("Failed to open the input file ... ", inFile)
		return
	}
	// 使用了defer关闭文件句柄
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line , seems unexpected ... ")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

/**
写入文件
*/
func writeValues(values []int, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create the output file ... ", outFile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

/**
冒泡排序
*/
func bubbleSort(values ...int) {
	flag := false
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}

/**
快速排序的实现
*/
func quitSort(left, right int, values ...int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quitSort(left, p-1, values...)
	}
	if right-p > 1 {
		quitSort(p+1, right, values...)
	}
}
