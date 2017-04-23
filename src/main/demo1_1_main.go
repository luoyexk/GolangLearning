package main

import (
	"math"
	"fmt"
	"runtime"
	 "time"
)



func main() {
	defer fmt.Println("这是延迟函数1")
	defer fmt.Println("这是延迟函数2")

	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	case "linux":
		fmt.Println("Linux.")
	}

	today := time.Now().Weekday()
	today2 := time.Now().Day()
	fmt.Println(today,today2)
	fmt.Println("测试延时函数 defer")
	fmt.Println("结尾")
}

func sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprint(math.Sqrt(-x), "这是转变后")
	}
	return fmt.Sprint(math.Sqrt(x))
}
