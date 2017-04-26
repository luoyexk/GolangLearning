package main

import (
	"os"
	"fmt"
)
/*
子网掩码demo有问题
 */
func main() {
	v1 := os.Args
	fmt.Println(len(v1)) //长度是1
	/*if len(os.Args) != 2 {
		//fmt.Fprintf(os.Stderr, "Usage: %s ip.addr\n", os.Args[0])
		os.Exit(1)
	}
	dotaddr := os.Args[1]
	addr := net.ParseIP(dotaddr)
	if addr == nil {
		fmt.Println("Invalid address")
	}
	mask :=addr.DefaultMask()
	fmt.Println("Subnet mask is:",mask.String())*/
}
