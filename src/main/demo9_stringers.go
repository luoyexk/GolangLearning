package main

import "fmt"

// 一个普遍存在的接口是 fmt 包中定义的 Stringer。
type Dog struct {
	name string
	age  int
}

func (d Dog) String() string {
	return fmt.Sprintf("这条狗的名字是：%s，今年%d岁了。", d.name, d.age)
}
func main() {
	dog1 := Dog{"大黄", 8}
	var dog2 = new(Dog)
	dog2.name = "二狗子"
	dog2.age = 2
	fmt.Println(dog1,dog2)
}
