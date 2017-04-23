package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

type KEY struct {
	p1 int
}

type VALUE struct {
	value1 int
	value2 string
}

var m1 = map[KEY]VALUE{KEY{1}: VALUE{1, "content"}}
var m = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

var m2 = map[string]int{"map1": 1, "map2": 2}
var m3 map[string]int

func main() {
	fmt.Println(m1)
	m3 = make(map[string]int)
	m3["test"] = 1
	//m2 = make(map[string]int)
	m2["test"] = 2
	fmt.Print(m2)
	a := m2["test"]
	fmt.Println("a=", a)
	delete(m2, "test")
	fmt.Println(m2)

	// map遍历
	m4 := map[string]int{"test1": 1, "test4": 4, "test2": 2}
	for key := range m4 {
		fmt.Println("key=", key, "\tvalue=", m4[key])
	}

	test3Value,ok := m4["test3"] // map获取key不存在的值
	fmt.Println(test3Value,ok)
}
