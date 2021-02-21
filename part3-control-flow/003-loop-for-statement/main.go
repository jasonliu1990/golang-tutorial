package main

import "fmt"

func main() {
	x := 1
	// for condition; like while
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done x")

	y := 1
	for {
		if y > 9{
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("done y")


}
