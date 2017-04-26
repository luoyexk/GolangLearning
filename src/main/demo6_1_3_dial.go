package main

import "net"

/*
Go对传统的Socket网络编程的过程进行了封装，无论期望使用寿命协议建立什么形式的连接，都只需要调用Dial()函数即可
该函数定义如下
func Dial(net,add string) (Conn, error)
Dial()支持tcp,tcp4,tcp6,udp,udp4,udp6,ip,ip4,ip6
建立成功后，可以使用conn对象的Write()成员方法，在接收数据时，可以使用conn对象的Read()成员方法
 */
func main() {
	var bytes []byte
	// 建立一个tcp连接
	connT, err := net.Dial("tcp", "192.168.1.1:4000")
	// 建立一个udp连接
	connU, err := net.Dial("udp", "192.168.1.2:5001")
	// 建立一个ICMP连接,使用协议名，域名为百度，省略端口号
	connIp4,err := net.Dial("ip4:icmp","www.baidu.com")
	// 建立一个ICMP连接，使用协议号，主机地址为"119.75.218.77"，省略端口号
	connIp4_,err :=net.Dial("ip4:1","119.75.218.77")
	if err!=nil {
		connT.Write(bytes)
		connU.Write(bytes)
		connIp4.Write(bytes)
		connIp4_.Write(bytes)
	}
}
