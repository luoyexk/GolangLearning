package main

import (
	"net"
	"fmt"
)

/*
在net包中，IP地址类型被定义成一个byte数组，即若干个8位组
IP地址分析，判断是否是ip地址

 */
func isIP(ipString string) {
	ip := net.ParseIP(ipString)
	if ip != nil {
		fmt.Printf("ip地址是：%v",ip)

	} else {
		fmt.Printf("%s不是IP地址",ipString)
	}
	s :=ip.String()
	fmt.Printf("将IP转换成字符串%s\n",s)
}

func main() {
	isIP("192.168.1.1")
	isIP("12:23:21:1:1:1:1:1")
	isIP("0:0:0:0:0:0:0:1")
}
