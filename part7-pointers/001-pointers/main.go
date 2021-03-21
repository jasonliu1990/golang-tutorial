package main

import "fmt"

func main(){
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // the address in memory

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	//var b *int = &a
	b := &a
	fmt.Println(b)
	fmt.Println(*b) // the value stored at an address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)


}
