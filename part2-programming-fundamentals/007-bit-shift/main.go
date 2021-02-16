package main

import "fmt"

const (
	_ = iota
	kb = 1 << (iota * 10) // iota = 1
	mb = 1 << (iota * 10) // iota = 2
	gb = 1 << (iota * 10) // iota = 3
)
func main()  {
	x := 4
	fmt.Printf("%d\t%b\n", x, x)
	// bit shift
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)

	// bit shift & const & iota
	// 比較複雜的用法; 作範例使用即可
	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
