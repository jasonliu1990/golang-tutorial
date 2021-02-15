package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var z int8 = -128

func main()  {
	x = 42
	y = 42.4242
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	//環境
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}