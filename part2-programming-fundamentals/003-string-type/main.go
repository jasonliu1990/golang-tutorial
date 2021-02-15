package main

import "fmt"
// strings are immutable,
// once created, it is impossible to change the contents of a string
func main()  {
	s := "hello world"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) //slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++{
		fmt.Printf("%#U ", s[i]) //UTF
	}
	fmt.Println("")
	for i, v := range s{
		//fmt.Println(i, v)
		fmt.Printf("index %d: hex %#x\n", i, v)
	}
}