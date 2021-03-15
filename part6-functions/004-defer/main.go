package main

import "fmt"

func main(){
	defer foo() //推遲
	bar()
}

func foo()  {
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}