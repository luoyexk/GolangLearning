package mathapp

import "fmt"
import (
	"math"
	"math/cmplx"
)

func main() {
	//1 import
	fmt.Println("Hello,world.", math.Pi)
	// 2 单返回
	fmt.Println(add(2, 3))
	// 3 多返回
	a, b := addString("hello", "world")
	fmt.Println(a, b)
	// 4 命名返回值
	fmt.Print("命名返回值")
	fmt.Println(split(10))
	// 5 变量
	var i int
	fmt.Println(i, c, pythor, java)
	// 6 变量
	fmt.Println(str1, str2, str3, x1, x2, x3, flag)

	fmt.Println(a_test, a_test_, a64, a64_)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// 常量 const
	const World = "这是常量"
	fmt.Println(World)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func addString(x, y string) (string, string) {
	return y, x
}

// 命名返回值，return没有参数，将返回各个变量的当前值
func split(sum int) (x, y int) {
	x = sum * 4 / 2
	y = x - 4
	return
}

var c, pythor, java bool
var str1 string = "str1"
var str2 = "str2"
var x1, x2 int = 1, 2
var str3, x3, flag = "str3", 1, true

var a_test int = 2147483647
var a_test_ int = -2147483648
var a16 int16 = 32767   //2^15-1
var a16_ int16 = -32768 //-2^15
var a64 int64 = 9223372036854775807
var a64_ int64 = -9223372036854775808
var uA uint = 4294967295
var maxInt uint64 = 1<<64 - 1
