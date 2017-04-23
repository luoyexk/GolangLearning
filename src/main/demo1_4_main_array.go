package main

import "fmt"

func main() {
	var array [4]int    // 定义一个长度为4的int类型数组
	fmt.Println(array)  // 打印结果[0 0 0 0]
	array[0] = 1        // 给第数组索引为1的元素赋值为1
	fmt.Println(array[0])  // 打印结果为1
	var flag = []bool{true,false,true}  // 定义一个长度为3的bool类型数组
	fmt.Println(flag)   // 打印结果为[true false true]

	s1 := []int{2,4,6,8}
	s2 := []int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(s1) // 内容为：[2 4 6 8]
	// 切片结构，type表示切片类型，length表示切片当前长度，capacity表示容量（可选参数），即使设定，也会在切片后被变更
	// var var_name []type = make([]type, length, capacity)
	var ss []int = make([]int,3,100)  // 声明一个长度为3，容量为100的切片
	// len(切片)表示切片的长度，cap(切片)表示切片最长可达到的容量
	fmt.Printf("len:%d\tcap:%d\tcontent:%v\n",len(ss),cap(ss),ss) // len:3	cap:100	content:[0 0 0]
	ss = s1[3:4] // 从s1中截取索引3~4（不含）的元素创建为一个新的切片
	fmt.Printf("len:%d\tcap:%d\tcontent:%v\n",len(ss),cap(ss),ss) // len:1	cap:1	content:[8]
	ss = s1[1:3] // 从s1中截取索引1~3（不含）的元素创建为一个新的切片
	fmt.Printf("len:%d\tcap:%d\tcontent:%v\n",len(ss),cap(ss),ss) // len:2	cap:3	content:[4 6]
	ss = s2[1:3] // 从s2中截取索引1~3（不含）的元素穿件为一个新的切片
	fmt.Printf("len:%d\tcap:%d\tcontent:%v\n",len(ss),cap(ss),ss) // len:2	cap:9	content:[2 3]
	// 由此可见，切片ss是一个长度可变，最大容量根据数组源的长度而改变的“动态数组”,切cap由源数组len-startIndex

}
