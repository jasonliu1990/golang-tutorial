package main

import "fmt"
// https://golang.org/pkg/fmt/
//％v按值的本來值輸出
//％+ v在％v基礎上，對結構體分段名和值進行展開
//％＃v輸出Go語言語法格式的值
//％T輸出Go語言語法格式的類型和值
//%%輸出％本體
//％b整型以二進制方式顯示
//％o整型以八二進制方式顯示
//％d整型以十二進制方式顯示
//％x整型以十六進制方式顯示
//％X整型以十六進制，字母大寫方式顯示
//％U Unicode字符
//％f浮點數
//％p指針，十六進制方式顯示
var y = 42
func main()  {
	fmt.Println(y)
	// type
	fmt.Printf("%T\n", y)
	fmt.Printf("%#x\n%b\n%x", y, y, y) // hex; binary
	// 格式化輸出
	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)
	fmt.Println("%v", y) // normal value
	// 格式化不同data type
	pi := 3.1415
	s2 := fmt.Sprintf("%v %v %v", "hello", pi, true)
	fmt.Println((s2)) // hello 3.1415 true
}
