package main

import "fmt"

// 定义一个结构体，命名为Parent
type Parent struct {
	mother string // 字段一为string类型
	father string // 字段二为string类型
}

// 定义一个结构体，命名为Me
type Me struct {
	age     int    // 字段一为int类型
	name    string // 字段二为string类型
	gender  string // 字段三为string类型
	parents Parent // 字段四为结构体Parent类型
}

func main() {
	var me Me
	me.age = 20
	me.name = "小李"
	me.gender = "男"
	me.parents.father = "李雷"
	me.parents.mother = "韩梅梅"
	fmt.Printf("age:%d \tname:%s \tgender:%s \tfather:%s \tmother:%s", me.age, me.name, me.gender, me.parents.father, me.parents.mother)
	fmt.Println()
	// 当然，你也可以这样赋值
	me2 := Me{38, "李雷", "男", Parent{"牛爱花", "李刚"}}
	fmt.Printf("age:%d \tname:%s \tgender:%s \tfather:%s \tmother:%s", me2.age, me2.name, me2.gender, me2.parents.father, me2.parents.mother)

	fmt.Println()
	//结构体指针
	var me_ptr *Me  // 什么结构体指针
	me_ptr = &me    // 把me的指针传递给me_ptr
	printMe(me_ptr) // 通过结构体指针获取到指针对应的结构体的字段
	fmt.Println(v1, p, v2, v3, v4)
}

func printMe(me *Me) {
	fmt.Println("我是：", me.name)
	fmt.Printf("爹是：%s\t娘是：%s\n", me.parents.father, me.parents.mother)
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
	v4 = Vertex{Y: 2, X: 4}
)
