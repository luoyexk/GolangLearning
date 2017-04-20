package main

import "fmt"
// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
// 1、定义一个人的接口，人的行为共性会说话
type IPerson interface {
	speak()
	//walk()
}
// 2、定义一个人，李雷
type LiLei struct {

}
// 3、写一个speak方法，方法接收者是type LiLei，说明LiLei实现了这个IPerson接口
func (l LiLei) speak()  {
	fmt.Println("Hi, I`m LiLei")
}

/*func (l LiLei) walk()  {
	fmt.Println("I can walk")
}*/

func main()  {
	var person IPerson // a、声明一个接口
	person = new(LiLei) // b、new一个李雷来实现person接口，如果我们没有写3这一步，我们是不能这样写的
	person.speak() // c、执行者其实是我们的LiLei，这很像Java中的多态
	// 打印结果：Hi! I`m LiLei
	// 4、此时在IPersion接口内再建立一个walk()方法，b处马上出现提醒：
	   // Cannot use new(LiLei) (type *LiLei) as type IPerson in assignment：不能通过new(LiLei)作为IPerson类型
	// 5、因为我们的type LiLei没有实现walk()方法，所以Go认为，LiLei没有实现IPerson接口
	// 6、只要我们补上一个方法，接收者为LiLei就可以了
	//person.walk()
}


