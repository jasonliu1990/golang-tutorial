package main

import "fmt"
// iota 特殊常量; 可被編輯器修改的常量
// 用在const時, 出現時重置為0, 每新增一行iota計數加1
// 可被用做枚舉值
const (
	a = iota
	b
	c
)
const (
	d = iota
	e
	f
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

