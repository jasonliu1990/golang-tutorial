package main

import "fmt"

func main()  {
	foo()
	// anonymous func
	//匿名func: func(){ <code> }()
	//記得加()
	func(){
		fmt.Println("anonymous func")
	}()
	// add args
	func(x int){
		fmt.Println("anonymous func", x)
	}(42)

	// func expression
	f := func(){
		fmt.Println("func expression")
	}
	f()
	f2 := func(x int){
		fmt.Println("func expression", x)
	}
	f2(42)

}

func foo(){
	fmt.Println("foo")
}
