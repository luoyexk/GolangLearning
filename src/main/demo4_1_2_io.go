package main

import (
	"os"
	"fmt"
	"strings"
	"io"
)

/*
  Go的IO接口一些简单读取函数
*/
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)   // 缓冲buffer，生成一个切片
	n, err := reader.Read(p) // 把数据读到切片中，n表示读了多少数据，err表示可能存在的错误
	if n > 0 {
		return p[:n], nil
	}

	return p, err // 为什么读取到n呢，因为buffer的长度可能大于我们给的字符的长度，所以通过切片的特性只返回我们读取到的位置

}

// 从字符串读取
func sampleReadFromString() {
	// 通过给定的字符串，建立一个读取器
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)
	fmt.Println(string(data))
}

// 从控制台输入读取字符串
func sampleReadFromStdin() {
	data, _ := ReadFrom(os.Stdin, 12)
	fmt.Println(data)         // 读取到的字节
	fmt.Println(string(data)) // 字节强转为字符查看我们到底读取到了什么
}

// 从文件读取字符串
func sampleReadFromFile(longth int) {
	file, _ := os.Open("d:/Go/workplace/GolangLearning/src/maindemo4_1_2_io.go")     // 打开我们的源码
	defer file.Close()                // 读取完毕后关闭文件流
	data, _ := ReadFrom(file, longth) // 读取我们的文件
	fmt.Println(data)                 // 打印数组切片
	fmt.Println(string(data))         // 打印数组切片强转字符串结果
}

func main() {
	//sampleReadFromString()
	  sampleReadFromStdin() // 编译build，然后用cmd打开该可执行文件，输入字符串即可看到从终端读取到的数据
	// D:\Go\workplace\dev\src\basiceio>basiceio.exe
	// 12345678901234567890num
	// [49 50 51 52 53 54 55 56 57 48 49 50]
	// 123456789012
	//sampleReadFromFile(9)
	//sampleReadFromFile(10)
}