package main

import "fmt"

var x int
func main(){
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
	//	閉包
	fmt.Println("closure")
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo(){
	fmt.Println("hello")
	x++
}

func incrementor() func() int{
	var x int
	return func() int{
		x++
		return x
	}
}