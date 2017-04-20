package main

import "fmt"

/**
方法
【方法接收者】 出现在 func 关键字和方法名之间的参数中。
 */
type Bird struct {
	name  string
	color string
}

func (b Bird) fly() {
	// 这里的方法接收者是Bird
	fmt.Println(b.color + "的" + b.name + "在飞")
}

type Num struct {
	a int
	b int
}

func (n Num) Add() {
	fmt.Println(n.a + n.b)
}

// 【方法接收者】为类型指针的方法
func (n *Num) ChangeNum() {
	// 计算后，n.a的值改变
	n.a = n.a * n.a
	n.b = n.b * n.b
}

// 【方法接收者】为类型的方法
func (n Num) ChangeNum2() {
	// 计算后，n.a的值不会改变
	n.a = n.a * n.a
	n.b = n.b * n.b
}
func main() {
	var bird = new(Bird)
	bird.name = "燕子"
	bird.color = "黑色"
	bird.fly()

	bird2 := &Bird{"喜鹊", "灰色"} //通过指针
	bird2.fly()

	n := Num{2, 4}
	n.Add()
	fmt.Println(n)
	n.ChangeNum()
	fmt.Println(n)
	n.ChangeNum2()
	fmt.Println(n)
}
