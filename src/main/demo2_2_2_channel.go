package main

import "fmt"

/*
缓冲通道
channel 可以是 带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
 */
func main() {
	c := make(chan string,2)
	c <- "我是1通道"
	c <- "我是2通道"
	//c <- "我是3通道" // 打开的话就发生all goroutines are asleep - deadlock!
	fmt.Println(<-c)
	fmt.Println(<-c)
}
