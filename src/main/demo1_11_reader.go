package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	var r io.Reader = strings.NewReader("hello world! I love Golang!") // 准备一个字符流
	//r2 := strings.NewReader("hello world! I love Golang!") // 这样写也是可以的
	bytes := make([]byte, 8) // 准备一个数组用于装载数据
	for { // 开始读取字符流中的数据
		m, err := r.Read(bytes) // 读取数据，返回读取到的数据的长度，和错误信息
		fmt.Printf("n = %v, err = %v, bytes = %v\n", m, err, bytes)
		fmt.Printf("bytes[:n]=%q\n", bytes[:m])
		if err == io.EOF {
			// 在遇到数据流结尾时，返回 io.EOF 错误
			break
		}
	}
}
