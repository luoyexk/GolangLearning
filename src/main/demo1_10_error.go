package main

import "fmt"
/**
error
Go 已经定义了一个error接口，其有一个Error()方法，我们只需要实现Error()这个方法即可
 */
type DivideError struct {
	a, b int
}

func (div DivideError) division() (result int, errorMsg string) {
	if div.b == 0 {
		errorMsg = div.Error()
	} else {
		result = div.a / div.b
	}
	return // 这里默认返回变量最后的值，如果变量没有被赋值则默认为空值，string空值为""
}
// DivideError实现了error接口
func (div DivideError) Error() string {
	return fmt.Sprintf("b:%d不能为零。", div.b)
}

func main() {
	d := DivideError{10, 5}
	if result, errorMsg := d.division(); errorMsg == "" {
		fmt.Println(result)
	}

	d2 := DivideError{10,0}
	if _,errorMsg :=d2.division();errorMsg!="" {
		// 因为division()是多返回值函数，但进入到此肯定发生了错误，所以我不在乎返回的result，使用“_”表示不关心结果
		fmt.Println(errorMsg)
	}
}
