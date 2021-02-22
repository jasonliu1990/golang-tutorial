package main

import "fmt"

// switch默認每個case後帶有break, 匹配成功後不會向下執行其他case並跳出switch
// 但可以使用fallthrough強制執行後面的case
func main() {
	fmt.Println("=====case1========")
	switch {
	case (2 == 4):
		fmt.Println("XXX") //不會print
	case (3 == 3):
		fmt.Println("print 1")
		fallthrough
	case (3 == 5):
		fmt.Println("not true 1")
		fallthrough
	case true:
		fmt.Println("print 2")
		fallthrough
	default:
		fmt.Println("this default")
	}
	fmt.Println("=====case2========")
	n := "bond"
	switch n {
	case "moneypenny":
		fmt.Println("money")
	// 多重條件
	case "bond", "Q":
		fmt.Println("bond or q")
	default:
		fmt.Println("not found")
	}

}
