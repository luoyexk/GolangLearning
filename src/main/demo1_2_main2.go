package main

import "fmt"
/**
指针的学习
指针表示：保存了变量的内存地址
&表示生成一个指向其作用对象的指针
*表示指针指向的底层的值
 */
func main()  {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)

	i, j := 42, 99
	var p *int= &i         // point to i // 把i的内存地址给p
	// 另一种写法 p := &i
	fmt.Println("p=",p,"    &i=",&i)//打印出p= 0x10b92140     &i= 0x10b92140
	fmt.Println("*p=",*p) // read i through the pointer // 读取p指针指向底层的值
	*p = 21         // set i through the pointer // 把21赋给p指针指向底层的值
	fmt.Println("i=",i)  // see the new value of i // 打印i的值

	p = &j         // point to j //把j的指针赋给p
	fmt.Println("p=",p,"   &j=",&j)
	*p = *p / 33   // divide j through the pointer // 获取指针p指向底层的值，然后除以33，把结果赋给底层的值
	fmt.Println("*p=",*p)
	fmt.Println("j=",j,"    &j=",&j) // see the new value of j // 因为底层的值被修改了，所以j的指针指向底层的值为3

	var ptr *int
	fmt.Printf("ptr:%x\n",ptr)
	fmt.Println("ptr:",ptr)
}
