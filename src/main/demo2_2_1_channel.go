package main

import "fmt"

/*
channel:一个有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
ch <- v    // 将 v 送入 channel ch。
v := <-ch  // 从 ch 接收，并且赋值给 v。
（“箭头”就是数据流的方向。）
使用channel前，必须通过make创建channel
ch := make(chan int)创建了一个int类型的管道
默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。
//?同一个管道变量，也遵循后进先出？
 */

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println("结果是：",sum)
	c <- sum // 将和送入 c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) // 结果是17
	go sum(a[len(a)/2:], c) // 结果是-5
	x, y := <-c, <-c // 从 c 中获取
	fmt.Println(x, y, x+y)
}


