package main

import "fmt"

func main() {
	fmt.Println("hello world")
	foo()
	fmt.Println("something...")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm foo")
}

func bar()  {
	fmt.Println("and then we exited")
}

// control flow:
// 1. sequence
// 2. loop; iterative
// 3. conditional
